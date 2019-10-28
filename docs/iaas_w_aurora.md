# IaaS with Aurora

You realized now that it's a difficult task to build a CloudFormation for Database.  
Also, the example that you deployed in the previous exercise is highly unsecure (as it's accessible from everywhere) and also not very reliable (as there is only one instance).  
Luckily, AWS provides a service called RDS (Relational Database Service). As part of RDS, there is a product called Aurora Serverless. It's a database that is completely compatible with MySQL and brings all the advantages of Serverless technology, means less resources to manage for you.
In this exercise we deploy now Aurora through CloudFormation and change the configuration of our Webserver to use the new database.  
In a real-world use case you would for sure migrate your existing data, but let's skip that for now.

## Step-by-Step instructions

1. Update the [cf-all.yaml](./files/cf-all.yaml) CloudFormation and add the cluster resource from  here: [cf-aurora.yaml](./files/cf-aurora.yaml).

1. Remove the Database resource and the according DatabaseSecurityGroup in the [cf-all.yaml](./files/cf-all.yaml), you don't need it anymore.

2. Commit and push your infrastructure source code changes to Github.

Congratulations, you have one EC2 instance less to manage! Also you get the following advantages:

- Transparent Scalability
- High availability
- Managed Backups

You can now move to the next chapter, where we will see how we can make our application easier transferable.