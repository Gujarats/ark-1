# Aws rotate keys 
You can use this project for rotate your secret/access keys using AWS Lambda function

# How to use
You need to set the environment variable in AWS lambda console.
And you must specify the prefix for all environment variable here is some example The prefix is `MYAPP`
your environment variable should looks like this

```shell
 MYAPP_REGION="YOUR_AWS_REGION"
 MYAPP_USER=YOUR_IAM_USER
 MYAPP_GROUP=YOUR_ARN_GROUP
```

And you must speciy the prefix in your environment variable like this

```shell
PREFIX_ENV_MINE=MYAPP
```
