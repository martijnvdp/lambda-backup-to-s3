package main

import (
	"github.com/CCV-Group/lambda-backup-rds2s3/pkg/handlers"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handlers.HandleRequest)
}
