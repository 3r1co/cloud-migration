# Infrastructure as a Service

In this first exercise, we are building an application consisting out of two components:

- A web server that is containg a so called CRUD application (Create, Read, Update, Delete)
- A database server that is hosting the applications data

## Step-by-Step instructions

1. Go to the AWS Console and deploy an EC2 Instance through the EC2 Console. This EC2 instance will act as your database. When going through the wizard, take care of the following things:

    - Choose the latest "Amazon Linux 2" AMI (Amazon Machine Image as image)
    - Choose "t2.micro" as instance type
    - In the security group screen, allow connections on port 3306 (MySQL) and 22 (SSH) from anywhere
    - When launching the instance, create a new key pair
        -   do not lose the file that is downloaded after creating the keypair, you need it to connect to your EC2 instance

1. After your EC2 instance is launched, connect to your instance with SSH:

    - Go in your terminal to the folder with the previously downloaded key pair
    - Connect with SSH like follows:

            ssh -i <your-key-pair.pem> ec2-user@<public-ip-of-your-ec2-instance>

2. Install MySQL on this EC2 instance:

        yum install -y mysql mysql-server

3. Run the following statement to create your database schema:

        wget https://3r1.co/cloud-migration/files/database.sql
        mysql -u root < database.sql

    With this, you have successfully installed and configured your database.  
    Now we move on the webserver installation.

1. Go again to the AWS Console and deploy an EC2 Instance through the EC2 Console. This EC2 instance will act as your webserver. When going through the wizard, take care of the following things:

    - Choose the latest "Amazon Linux 2" AMI (Amazon Machine Image as image)
    - Choose "t2.micro" as instance type
    - In the security group screen, allow connections on port 8080 (HTTP) and 22 (SSH) from anywhere
    - When launching the instance, use your existing key pair.

1. After your EC2 instance is launched, connect to your instance with SSH:

    - Go in your terminal to the folder with the previously downloaded key pair
    - Connect with SSH like follows:

            ssh -i <your-key-pair.pem> ec2-user@<public-ip-of-your-ec2-instance>

1. Download and start the sample webserver application:

        wget https://3r1co-github-artifacts.s3.amazonaws.com/crud-app.tar.gz
        tar xzvf crud-app.tar.gz
        cd crud-app
        ./crud-app --dbHost <private-ip-of-your-database-instance> --dbName isen --dbUser isenuser --dbPass isen1234

1. In your browser go to http://<public-ip-of-your-ec2-instance>:8080 .  
You should see your application running.

Congratulations, you have now deployed our first application based on VM Instances.

## Follow up exercise

To get familiar with CloudFormation, there is a sample file available [here](./files/cf-application.yaml).  

1. Deploy this file with CloudFormation.  

1. Currently, the cf-application.yaml is deploying a single EC2 Instance. Your task is now to transform the EC2 Resource into a [Launch Configuration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html). Afterwards, assign this Launch Configuration to an [Autoscaling Group](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html). 

## Next steps 

You can now move to the next chapter, where we will automate the deployment of our application.