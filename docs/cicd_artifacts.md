# Continuous Integration and Delivery

In an enterprise context, it's normally not one person doing the build, test and deploy steps for an application. You rather rely on an automated system which gives you clear visibility on _who_ did _when_ _which_ change.  
In this lab, we will not build a very advanced CI/CD system as you have a dedicated course for that, but we will try to focus on a MVP (minimum viable product) configuration.

At the end of this lab, every commit to the master branch of your repository should result in a build process to start.

## Step-by-Step instructions

1. Create a Repository on github.com in which you will store  

    - your application source code
    - your infrastructure source code

1. Build artifact
1. Upload to S3
1. Deploy with CloudFormation API