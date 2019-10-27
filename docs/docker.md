# Containerisation

In this tutorial, we would like to make our application easily transferable between different cloud providers.  
As you have learned throughout the course, containerisation is a good way to achieve that goal.  
In this tutorial, we will add a Dockerfile to our repository and push it to docker.io. This enables us to have our artifacts stored in a "neutral" place, which can be used by various cloud providers.  
Once this is done, we will address the deployment of container through ECS, the Container service of AWS.