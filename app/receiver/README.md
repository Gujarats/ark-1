# Receiver 
This app is use for getting the access & secret keys from parameter store and store it localy to : 
 - gradle.properties
 - environment variables

You can choose one of them or both !

## How to install

You can choose the binary file ready to download on this [release](Gujarats/ark/releases).
Or you can run this command : 

``` shell
$ curl https://raw.githubusercontent.com/Gujarats/ark/master/app/receiver/downloader.sh | sh
```

Finally you must config your receiver binary and the `Configuration` section below.

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

After above steps the app is ready to use and the binary is in the `./bin` directory. And now you need to test it by running this commands :

```shell
$ cd bin
$ ./receiver
```

It should print the secret & access keys if your set your configuration correctly.
