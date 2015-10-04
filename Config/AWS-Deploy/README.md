**DETER-BISTRO TRAVIS CI -- SOFTWARE BUILD DEPENDENCY MANAGEMENT AND TESTING**  
DeterBistro uses the Travis CI continuous integration service to manage software, framework and package dependencies for build management and target configurations.  
  
Travis CI automatically detects when a commit has been made and pushed to a GitHub repository that is using Travis CI, and each time this happens, it will try to build the project and run tests. This includes commits to all branches, not just to the master branch. Travis CI will also build and run pull requests. When that process has completed, it will notify a developer in the way it has been configured to do so[3] — for example, by sending an email containing the test results (showing success or failure), or by posting a message on an IRC channel. In the case of pull requests, the pull request will be annotated with the success/failure and a link to the build log, using a GitHub integration.

**DETER-BISTRO AWS CODEDEPLOY -- SOURCE CODE DEPLOYMENTS**  
AWS CodeDeploy coordinates application deployments to Amazon EC2 instances, on-premise instances, or both. (On-premise instances are physical devices that are not Amazon EC2 instances.)  
  
An application can contain deployable content like code, web, and configuration files, executables, packages, scripts, and so on. AWS CodeDeploy deploys applications from Amazon S3 buckets and GitHub repositories.  
  
Define an appspec file which basically is just a yml file that lays out the list of shell scripts that get run to deploy your app. Travis CI packages up appspec.yml and DeterBistro source code into zip file from GitHub repo and pushes it into AWS EC2 instance. Define in CodeDeploy app spec.yml DeterBistro source files/directories and GitHub repo. Define in CodeDeploy EC2 instances.
