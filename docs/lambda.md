# All in with Serverless

The ultimate goal of application refactoring is making it Serverless.
The application in the `/src` directory is already compatible with AWS Lambda.
You now have to find out how to: 

- Compile a Golang application for AWS Lambda
- Package the application to be compatible with AWS Lambda (ZIP Format)
- Deploy the Lambda package

Additionally, create an Application Load Balancer to routes traffic to your Lambda function.

## Follow-up exercise

Create a Github Workflow that automatically builds, packages and deploys your Lambda Function.