# Containerisation

In this tutorial, we would like to make our application easily transferable between different cloud providers.  
As you have learned throughout the course, containerisation is a good way to achieve that goal.  
In this tutorial, we will add a Dockerfile to our repository and push it to AWS ECR. 
Once this is done, we will address the deployment of container through ECS, the Container service of AWS.

There is already a Dockerfile in the `/src` directory of this repository. You can use it to build the Docker image on your local workstation.

For this tutorial, you'll have to use the AWS Documentation and find out the following things:

- How to log in with the AWS CLI to AWS ECR
- How to push a Docker image to AWS ECR
- How to deploy a Docker image to AWS ECS

## Follow-up exercise

Check the pipeline in the `.github/workflows/docker.yaml`. Make the pipeline work so that a new push to your Github repository results in a new Image Build + Push + Deployment.