# AwsLambdaHttpEndpoint

This is an example AWS lambda function written in GO, to be used as HTTP endpoint.
Notice that you need to upload to AWS a compiled GO Application, not just GO source code!
see: https://yos.io/2018/02/08/getting-started-with-serverless-go/

# Building deployment package

```sh
go build hello.go
zip hello.zip hello
```

# Deploying to AWS lambda

Just create new AWS lambda function for GO and upload the ZIP file