package utils

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

func AwsErrorPrint(err error) {
	if aerr, ok := err.(awserr.Error); ok {
		fmt.Println(aerr.Error())
	} else {
		fmt.Println(err.Error())
	}
	fmt.Println(os.Stderr)
	os.Exit(1)
}

func SlackSend(task string) {
	name := "Go"

	jsonStr := `{"username":"` + name + `","text":"/remind me to ` + task + ` in 3 hours"}`

	req, err := http.NewRequest(
		"POST",
		"https://hooks.slack.com/services/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",
		bytes.NewBuffer([]byte(jsonStr)),
	)

	if err != nil {
		fmt.Print(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(resp)
	defer resp.Body.Close()
}
