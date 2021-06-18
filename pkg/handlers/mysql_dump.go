package handlers

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"time"
)

func run_mysql_dump(host, database, port, username, password, bucket string) (err error) {
	var out bytes.Buffer
	cmd := exec.Command("mysqldump",
		"--single-transaction=TRUE",
		"-h", host,
		"-p", port,
		"-u", username,
		fmt.Sprintf("-p'%s'", password),
		database,
		"| gzip -c | aws s3 cp - s3://", bucket,
		"/", database,
		"_", time.Now().Format("20060102150405"))

	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%q\n", out.String())
	if err != nil {
		return err
	}
	return nil
}
