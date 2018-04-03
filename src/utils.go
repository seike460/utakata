package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

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
	v := url.Values{}
	v.Set("token", "-")
	v.Add("time", strconv.FormatInt(time.Now().Unix()+600, 10))
	v.Add("text", task)
	fmt.Println(v.Encode())
	url := "https://slack.com/api/reminders.add?" + v.Encode()
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray)) // htmlをstringで取得
}
