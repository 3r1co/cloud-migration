# Continuous Integration and Delivery

In an enterprise context, it's normally not one person doing the build, test and deploy steps for an application. You rather rely on an automated system which gives you clear visibility on _who_ did _when_ _which_ change.  
In this lab, we will not build a very advanced CI/CD system as you have a dedicated course for that, but we will try to focus on a MVP (minimum viable product) configuration.

At the end of this lab, every commit to the master branch of your repository should result in a build process to start.

## Step-by-Step instructions

1. Fork [this](https://github.com/3r1co/cloud-migration) repository. It contains   

    - the application source code (in src/)
    - the infrastructure source code (in docs/files/)

1. Create a programmatic IAM user in AWS.  
    -  Give EC2FullAccess and CloudFormationFullAccess managed policy to this user.
    -  Generate an Access Key ID and Secret Access Key bundle.
1. Add the following secrets to your previously forked Github repository:

    - AWS_ACCESS_KEY_ID : from the previously created bundle.
    - AWS_SECRET_ACCESS_KEY : from the previously created bundle.
    - DBName : choose freely
    - DBPassword : choose freely
    - DBRootPassword : choose freely
    - DBUser : choose freely
    - KeyName : the name of the keypair you created in the previous exercise

1. Create an S3 bucket that will be used to store your artifacts. Uncheck the box asking you to forbid making the bucket public. Add a bucket policy allowing the previously created user to upload artifacts and change an objects ACL (Access Control List).

