package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Database struct {
	Database        string `json:"database"`
	Host            string `json:"host"`
	Port            string `json:"port"`
	Username        string `json:"username"`
	Aws_sm_item     string `json:"aws_sm_item"`
	Aws_sm_item_key string `json:"aws_sm_item_key"`
}

var TMP string

func init() {
	var err error
	TMP, err = ioutil.TempDir("", "")
	if err != nil {
		log.Fatal(err)
	}
}

func HandleRequest(ctx context.Context, databases []Database) {
	var db_password map[string]string
	AWS_Region := os.Getenv("AWS_REGION")
	AWS_Bucket := os.Getenv("BUCKET_NAME")
	AWS_Account := os.Getenv("AWS_ACCOUNT_ID")

	if AWS_Region == "" || AWS_Bucket == "" || AWS_Account == "" || databases == nil {
		fmt.Println("missing variables")
		os.Exit(2)
	}

	for _, i := range databases {
		secret, err := GetAWSSecretValue(i.Aws_sm_item)
		json.Unmarshal([]byte(secret), &db_password)
		if err != nil {
			fmt.Println("Password could retrieved from secretsmanager")
			os.Exit(2)
		}
		err = run_mysql_dump(i.Host, i.Database, i.Port, i.Username, db_password[i.Aws_sm_item_key], AWS_Bucket)
		if err != nil {
			fmt.Println("dump failed")
			os.Exit(2)
		}
	}

	defer os.RemoveAll(TMP) // clean up
}
