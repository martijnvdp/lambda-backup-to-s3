package handlers

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func GetAWSSecretValue(id string) (secretstring string, err error) {
	s := session.Must(session.NewSession())
	sm := secretsmanager.New(s)
	output, err := sm.GetSecretValue(&secretsmanager.GetSecretValueInput{SecretId: &id})
	if err != nil {
		panic(err.Error())
	}
	return *output.SecretString, err
}
