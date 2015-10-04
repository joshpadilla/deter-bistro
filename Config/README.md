
**DETERBISTRO BUILD, DEPLOY AND TEST INFRASTRUCTURE**  
Summary of components built and DevOps OSS tools chosen. There are also detailed readme(s) in each section of the GitHub DeterBistro **"deter-bistro/Config"** directory.  
  
1) Upstream Dependency Management, Built In Upgrade Paths, Automated Testing (Travis CI YAML)  
-Managed via ".travis.yml" build script (https://github.com/joshpadilla/deter-bistro/blob/master/.travis.yml)  
-Real time build status at: https://travis-ci.org/joshpadilla/deter-bistro)  
  
2) Source Code Deployment Management (GitHub CodeDeploy Webhook & AWS YAML CodeDeploy)  
-Managed via "appspec.yml" deploy script (https://github.com/joshpadilla/deter-bistro/blob/master/appspec.yml)  
-Managed via "start.sh" and "stop.sh" deploy scripts  
 (https://github.com/joshpadilla/deter-bistro/tree/master/Config/AWS-Deploy)  
-Detailed ReadMe at (https://github.com/joshpadilla/deter-bistro/blob/master/Config/AWS-Deploy/README.md)  
  
3) Source Code/Repo Revision Control Management (GitHub)
    -Managed via Public GitHub Repo (
https://github.com/joshpadilla/deter-bistro)
    -Detailed Repo Overview ReadMe at (
https://github.com/joshpadilla/deter-bistro/blob/master/README.md)

4) AWS Network/Security Infrastructure (AWS CloudFormation JSON Template)
    -Managed via "AWS-Network.json" CloudFormation Template Script (
https://github.com/joshpadilla/deter-bistro/blob/master/Config/AWS-Network/AWS-Network.json
)
    -Includes implementations for AWS Key Pairs, IAM Roles, Security
Groups, VPC, Public/Private Subnets, NAT, Route Tables, Internet Gateways
    -Detailed ReadMe at (
https://github.com/joshpadilla/deter-bistro/blob/master/Config/AWS-Network/README.md
)

5) AWS Server/EC2 Instance Infrastructure (AWS CloudFormation JSON Template)
    -Managed via "AWS-Server.json" CloudFormation Template Script (
https://github.com/joshpadilla/deter-bistro/blob/master/Config/AWS-Network/AWS-Server.json
)
    -Includes implementations for EC2 Instance Region, Avail Zones,
Instance Type, Instance Role/Profile, OS Type, OS Default Packages, OS
Default Accts, SSH Keys
    -Detailed ReadMe at (
https://github.com/joshpadilla/deter-bistro/blob/master/Config/AWS-Server/README.md
)

6) AWS RDS MySQL & PostgreSQL Database Infrastructure (AWS CloudFormation
JSON Template)
    -Managed via "AWS-RDS-MySQL.json" & "AWS-RDS-PostgresSQL.json"
CloudFormation Template Scripts
    -(
https://github.com/joshpadilla/deter-bistro/blob/master/Config/AWS-Database/AWS-RDS-MySQL.json
)
    -(
https://github.com/joshpadilla/deter-bistro/blob/master/Config/AWS-Database/AWS-RDS-PostgreSQL.json
)
    -Includes implementations for RDS DB Engine, Version Number, Allocated
Storage, DB Names, Usernames, Passwords
    -Detailed ReadMe at (
https://github.com/joshpadilla/deter-bistro/blob/master/Config/AWS-Database/README.md
)
