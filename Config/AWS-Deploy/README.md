DeterBistro App Deployments with AWS CodeDeploy


Overview
AWS CodeDeploy is an automated way to manage the deployment of a software application across a complex EC2 deployment set using a simple application specification file within the source code of your project. You can define separate deployment environments within a single project, each can manage any number of resources. CodeDeploy provides you a way to fully automate the application deployment process onto EC2.

What is AWS CodeDeploy?
AWS CodeDeploy is a service provided by Amazon that allows the creation of workflows that automate the rollout of software applications to sets of EC2 instances. This is done by defining applications and groups of instances together and providing a simple programmatic application specification (AppSpec) to manage the deployment process details.

Prerequisites
This document expects that you have the following:
An active AWS account for testing with.
A moderate understanding of EC2.
An active EC2 key pair.
Familiarity with the AWS command line tools.
Working version of the AWS command line tools.
Core Concepts
Applications
An application is the main container of a CodeDeploy workflow. Every different software application you intend to deploy using CodeDeploy should be its own application within CodeDeploy. This allows you to manage deployment groups and revisions independently for each application you are managing. An application contains deployment groups and deployment groups have revisions deployed to them.
AppSpec
An AppSpec is an special YAML based file that describes the deployment process of your applications artifact to an EC2 instance. An AppSpec file must be included within every artifact that is deployed via CodeDeploy as a revision of an application.
Artifact
A specific build of your source applications. This could be a simple as a zip archive of application files containing an AppSpec file or the result of a complex build process (i.e. Java war, etc). CodeDeploy also offers integration with GitHub, this allows an artifact to also be a specific revision within a Git repository.
Deployment Configuration
A deployment configuration is set of rules tied to a deployment group on how to go about deploying a revision of an artifact to the set of EC2 instances within the deployment group. It has rules about how many machines to upgrade at once (i.e. one at time, half, etc) as well as rules on how to judge a revision deployment successful (or not).
Deployment Group
A deployment group is how CodeDeploy groups EC2 instances into pools for deploying a specific revision of an artifact. Each deployment group finds EC2 instances by a specific tag on the EC2 instance. Also a deployment group has a deployment configuration which depicts how a revision gets rolled out to the set of EC2 instances. A deployment group is synonymous with the concept of a deployment environment (i.e. development, test, staging, production).
Instance (AWS EC2 Server)
An instance is simply an EC2 instance configured to be managed by CodeDeploy. CodeDeploy will not create new instances on the fly (though if used with auto scaling rule, it would be possible). Instances must be configured in a special way in order to work with CodeDeploy. Any Amazon Linux AMI will work out of the box; other AMIs will need the CodeDeploy agent installed and configured.
Revision
A specific revision of your application as tracked by CodeDeploy. This is tied to an artifact of your source application. This artifact is usually a revision within a Git repository (posted on GitHub).
CodeDeploy Workflow Overview
CodeDeploy at its core is a software deployment workflow for EC2 instances. It is designed to allow you to automate the process of updating AWS EC2 instance servers from one release of your software to the next.

A CodeDeploy application contains deployment groups and deployment groups contain instances.

In order to do a deployment of your application you must first create an archive of your applications release artifact, pair it with an AppSpec file and upload it into GitHub. Once the software you intend to deploy is ready you initiate a deployment by providing the uploaded revision archive to the deployment group you wish to deploy onto.

Deployments can be initiated upon each commit into GitHub.
Configuring an Application in CodeDeploy
In order to deploy an application to EC2 using CodeDeploy you must first register and configure the application from within CodeDeploy itself before you can execute your first deployment. All of these steps below can be completed from the AWS Console, however we will (mostly) look at the CLI commands here.
Create an IAM Role for CodeDeploy
Since CodeDeploy will be calling AWS commands on your behalf, you will have to create a role for CodeDeploy inside of AWS IAM in order for CodeDeploy to work correctly. Frankly it is simpler to do this using the AWS Console than it is to use the CLI. Login to the AWS console IAM roles page and then create a role following the instructions from AWS here. Keep the roles ARN for later use.
Create the Application
The second step in creating a CodeDeploy workflow is to create an application entry inside of AWS CodeDeploy. The command is as follows:

aws deploy create-application \
  --application-name DemoApp

The CodeDeploy application allows you to manage separate workflows for different applications on the same AWS account. For example if your company maintained more than one product on AWS, creating multiple CodeDeploy applications, one for each product would be a way to isolate the workflows from one another.
Create the Development Group
The third step in creating a CodeDeploy workflow is to create a deployment group within an application. Remember the deployment group is used to manage which EC2 instances get deployed onto when executing a deployment. Identifying EC2 instances that should be managed by the deployment group is done by specifying a metadata tag that is set on each instance. The command is as follows:

aws deploy create-deployment-group \
  --application-name DemoApp \
  --deployment-group-name Production \
  --deployment-config-name CodeDeployDefault.OneAtATime \
  --ec2-tag-filters Key=CDG,Value=DemoApp_Production,Type=KEY_AND_VALUE \
  --service-role-arn arn:aws:iam::EXAMPLE:role/CodeDeploy

Create an Instance Profile
A special instance profile must be created in order for instances that are in a development group to execute a deployment. Each instance will need to be able to pull files from S3, read CodeDeploy configurations, etc.

Spin up EC2 Instances as Deployment Targets
Finally it is time to spin up a few instances that will fall into the development group as a deployment target. For simplicity we will use Amazon Linux AMIs as the target. However it is possible to use CodeDeploy for any number of other AMIs.

To spin up two instances use the command below:

aws ec2 run-instances \
  --image-id ami-b66ed3de \
  --key-name cli \
  --instance-type t2.micro \
  --iam-instance-profile Name=CodeDeployInstanceRole \
  --count 2

With the instances created we must now add the tag in order for those instances to be recognized by the development groups within your application in CodeDeploy. The command is as follows:

aws ec2 create-tags \
  --resources i-043938ee i-053938ef \
  --tags Key=CDG,Value=DemoApp_Production

Install CodeDeploy Agent on Instances
In order to push applications using CodeDeploy you must have the CodeDeploy agent installed and running. This agent does NOT come preinstalled even on Amazon Linux. To do this log into your instances run the following commands:

sudo yum -y update
sudo yum install -y aws-cli
cd /home/ec2-user
aws s3 cp s3://aws-codedeploy-us-east-1/latest/install . \
  --region us-east-1
chmod +x ./install
sudo ./install auto

This will install and start the CodeDeploy agent, ready to accept deployment requests.

Now we have a CodeDeploy application workflow defined with a development group and a few instances. We can now move onto understanding how to configure a software application to deploy and update instances via CodeDeploy.
Configuring your Software Application to Deploy Using CodeDeploy

In order for CodeDeploy to be able to understand how to deploy a software application you must define an application specification to accompany the software artifact at deployment time. This is done by creating an AppSpec file (appspec.yml) and placing it within the deployment artifact.

An AppSpec is a YAML based file that describes actions for CodeDeploy to take on each instance as it executes the deployment. The basic order of operations in a deployment is as follows:

A deployment is initiated targeting a specific deployment group and revision.
Each instance in the deployment group does the following (depending on the details of the deployment configuration):
The artifact from the revision is transferred to the instance from S3 (or GitHub).
If needed the artifact is unarchived.
The appspec.yml file is read from the root of the artifact and all steps in the AppSpec are executed.
As long as all steps complete without error the instance is considered deployed successfully and process for the instance is complete.
Once all instances are deployed the process is complete and the deployment group is updated to reflect that it is now synced to the new revision.
AppSpec File Format
In the deployment workflow described above, the AppSpec is a critical component in controlling a deployment. Therefore this file lets you describe a (somewhat) complex set of steps for the installation process of your software application.

The basic structure of an AppSpec file is as follows:

version: 0.0
os: operating-system-name
files:
  source-destination-files-mappings
permissions:
  permissions-specifications
hooks:
  deployment-lifecycle-event-mappings
version
The version of the AppSpec; should be left as 0.0.
os
The operating system this AppSpec is meant for; either 'linux' or 'windows'.
files
A mapping of files to copy from the revision artifact and where to put them on the instance. This is the core of the installation process (copying the files into place). The basic mapping is as follows:

files:
   - source: source-file-location
 	destination: destination-file-location

You can have as many source and destination mappings as needed.
permissions
A mapping of permissions to set on the files described in in the files section.
hooks
Hooks are a way to inject additional steps into the deployment workflow while CodeDeploy is running a deployment on an instance. There are hooks for things like stopping the application before the new files are copied, pre and post install and validating that the application is back up after the process is complete.
AppSpec Example
appspec.yml file

version: 0.0
os: linux
files:
  - source: /
	destination: /var/www/html/WordPress
hooks:
  BeforeInstall:
	- location: scripts/install_dependencies.sh
  	timeout: 300
  	runas: root
  AfterInstall:
	- location: scripts/change_permissions.sh
  	timeout: 300
  	runas: root
  ApplicationStart:
	- location: scripts/start_server.sh
  	timeout: 300
  	runas: root
  ApplicationStop:
	- location: scripts/stop_server.sh
  	timeout: 300
  	runas: root

A full reference for the AppSpec file can be found here.
Executing a Deployment
Now that we have CodeDeploy setup, a few instances created and an application configured; the final step is to execute a deployment.
Push Artifact to GitHub
To trigger a deployment, the artifact must be pushed up to GitHub for deployment.

If you do not have an S3 bucket you can create one with the following command:

aws s3api create-bucket \
  --bucket jmcodedeploy

Once you have a bucket the artifact can be pushed with the following command:

aws deploy push \
  --application-name DemoApp \
  --s3-location s3://jmcodedeploy/v001

The output of this command will give a template for the command needed to trigger the deployment.
Trigger Deployment
To trigger the deployment you can use the following command:

aws deploy create-deployment \
  --application-name DemoApp \
  --s3-location bucket=jmcodedeploy,key=v001,bundleType=zip,eTag="2145cafc7df9a4d1c22b8ffe5c35c9bb" \
  --deployment-group-name Production \
  --deployment-config-name CodeDeployDefault.OneAtATime \
  --description "Push of the application."

Demo Application
All of the commands and steps are covered in detail in the sections above; therefore we will not repeat all commands and details here. Instead we will list steps in order and focus on the demo details specifically.
Prerequisites
In order to work with the demo application you must first clone or download the code 
You must install and configure the AWS CLI tools.
Setup CodeDeploy and Deploy a Few Instances
Ensure you have your service role for CodeDeploy created and the ARN handy.
Use the command line or the AWS console to create and application in CodeDeploy.
Create two deployment groups. Let's call these 'Staging' and 'Production' This will help us demonstrate rolling an application out into several environments. In the command example above, be sure to change the tags to differentiate between staging and production. You don't want the same EC2 instances falling into one group.
Ensure you have an instance profile setup for you CodeDeploy managed EC2 instances.
SETUP SECURITY GROUPS: Make sure ports 22 and 8080 are in your 'default' security group (or whichever group you intend to use). If these are not correct you will not be able to connect to these instances.
Create instances. Setup three instances with a command similar to the following (ensure key-name is correct):

aws ec2 run-instances \
  --image-id ami-b66ed3de \
  --key-name cli \
  --instance-type t2.micro \
  --iam-instance-profile Name=CodeDeployInstanceRole \
  --count 3

We will use one instance as 'Staging' and the other two for 'Production'. Set the tags for the three instances with a commands similar to the following:

aws ec2 create-tags \
  --resources i-fa3a2210 \
  --tags Key=CDG,Value=DemoApp_Staging

aws ec2 create-tags \
  --resources i-f83a2212 i-f93a2213 \
  --tags Key=CDG,Value=DemoApp_Production

SSH into each instance and install the CodeDeploy agent. You could also do this on a single instance and then turn that instance into an AMI for reuse if you like. The user should be ec2-user.
We are now ready to deploy applications using CodeDeploy.
Trigger a Deployment and View the Results
Create an S3 bucket to hold your artifacts.
Open a terminal on your local machine, change directory (cd) to the root of the demo source code.
Use the push command to prepare the CodeDeploy artifact and upload it to S3.
To simulate a multiple environment deployment we will push to our 'Staging' deployment group first and test everything before we deploy to 'Production'. Use a command similar to the following to kick off the deployment using CodeDeploy:

aws deploy create-deployment \
  --application-name DemoApp \
  --s3-location bucket=jmcodedeploy,key=v001,bundleType=zip,eTag="4aded3cfd6780bf873c3079df579f2a3" \
  --deployment-group-name Staging \
  --deployment-config-name CodeDeployDefault.OneAtATime \
  --description "Push of the application to staging."

Now that we have deployed to 'Staging', first ensure the deployment worked by checking the status in the CodeDeploy in the AWS Console.
Now lets open the staging application in the browser. You should be able to load: http://[TARGET INSTANCE IP]:8080 and http://[TARGET INSTANCE IP]:8080/version.
If everything is working like it should you can then push the same artifact to the 'Production' deployment group with a command similar to the following:

aws deploy create-deployment \
  --application-name DemoApp \
  --s3-location bucket=jmcodedeploy,key=v001,bundleType=zip,eTag="4aded3cfd6780bf873c3079df579f2a3" \
  --deployment-group-name Production \
  --deployment-config-name CodeDeployDefault.OneAtATime \
  --description "Push of the application to production."

Now you should be able to load: http://[TARGET INSTANCE IP]:8080 and http://[TARGET INSTANCE IP]:8080/version on both production EC2 instances.
Congratulations you have now deployed a web application to multiple groups of EC2 instances using CodeDeploy.


