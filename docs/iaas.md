# Infrastructure as a Service

In this first exercise, we are building an application consisting out of two components:

- A web server that is containg a so called CRUD application (Create, Read, Update, Delete)
- A database server that is hosting the applications data

## Step-by-Step instructions

1. Go to the AWS Console and deploy the follow CloudFormation Stack:

    - [cf-database.yaml](../files/cf-database.yaml)  
    **Note the Private IP Address of the EC2 Instance you have just created. You will need it for the next step when you deploy your actual application.**

2. In CloudFormation Console, deploy now the stack for your application:

    - [cf-application.yaml](../files/cf-application.yaml)  
    **This stack outputs a public IP Address. Open it in your browser and you should see your deployed application.**


Congratulations, you have now deployed our first application based on VM Instances.

## Follow up exercise

Currently, the cf-application.yaml is deploying a single EC2 Instance. Your task is now to transform the EC2 Resource into a [Launch Configuration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html). Afterwards, assign this Launch Configuration to an [Autoscaling Group](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html). 

## Next steps 

You can now move to the next chapter, where we will automate the deployment of our application.