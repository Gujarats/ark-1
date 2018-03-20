# Aws rotate keys  [![Build Status](https://secure.travis-ci.org/Gujarats/ark.png)](http://travis-ci.org/Gujarats/ark)
A very simple script to rotate access & secret keys in AWS.
Please see the `app` folder to see more application command.

This script do theese : 
 - use Lambda function to rotate the secret & access keys and store it to parameter store
 - get the latest and active secret & access keys from parameter store  and save it local 
    * `$HOME/.gradle/gradle.properties`
    * Or you can choose environment variable to store it

 To see more details go to the `app` and read the `README.md` file for each application.



