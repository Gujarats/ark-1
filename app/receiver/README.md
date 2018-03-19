# Receiver 
This app is use for getting the access & secret keys from parameter store and store it localy to : 
 - gradle.properties
 - environment variables

You can choose one of them or both !

## How to install

Make sure you have setup Go in your environtment

``` shell
$ go get github.com/Gujarats/ark
$ cd $GOPATH/src/github.com/Gujarats/ark/app/receiver
$ go install
```

Basicly the app is already installed and ready to use, but for the env variable there is bit tricky to use this app.
Please follow this required steps : 

```shell
$ touch my-udpate-key
$ chmod +x my-update-key
$ vim my-update-key

#!/bin/bash
exec receiver
```

Add bash script to execute the app to avoid having a shell in a [shell](https://stackoverflow.com/questions/17368392/environment-variable-is-not-set-on-terminal-session-after-setting-it-with-os-p)

## Configuration

The configuration file is default location in `.ark/config.yaml` you can choose different type of config extension like JSON, YAML, TOML. Here is some config example using `.yaml`: 

```yaml
---
region : your-region 
profile : your-profile 
roleName : your-role-arn 
secretKey : ask your administrator to get the value from parameter store using defined key
accessKey : ask your administrator to get the value from parameter store using defined key
gradleAccessKey: ask your administrator what key to use 
gradleSecretKey: ask your administrator what key to use 
useGradleProperties : false
useEnvVariable : true

```

The rest of the config you can see in the `config.go` see what are the keys available
