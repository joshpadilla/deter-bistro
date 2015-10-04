**AWS MySQL & PostgreSQL Databases**  
The DeterBistro MySQL & PostgreSQL databases are provisioned, modified and managed via AWS CloudFormation JSON files. They support parameter definitions, resource (database instance servers) creation, modification and deletion. Each of the database engines are run as distributed relational database services.

The DeterBistro MySQL database is accessible at a canonical name (CNAME) and address of “bistrofridge.deter-bistro.net:3306” It has a service address of “bistrofridge.cccchoqs3xpb.us-west-1.rds.amazonaws.com:3306” and has 5GB of allocated storage. The MySQL database engine is currently at version 5.6.

The DeterBistro PostgreSQL database is accessible at a canonical name (CNAME) and address of “kitchenfridge.deter-bistro.net:5432” It has a service address and name of “kitchenfridge.cccchoqs3xpb.us-west-1.rds.amazonaws.com:5432” and has 5GB of allocated storage. The PostgreSQL database engine is currently at version 9.4.
