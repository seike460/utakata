package main

import (
	"fmt"
	"os"

	"github.com/faultline/faultline-go/faultline"
	"github.com/seike460/utakata/src/utils"
)

var notifications = []interface{}{
	faultline.Slack{
		Type:           "slack",
		Endpoint:       os.Getenv("FAULTLINE_NOTIFY_SLACK_ENDPOINT"),
		Channel:        os.Getenv("FAULTLINE_NOTIFY_SLACK_CHANNEL"),
		Username:       "faultline-notify",
		NotifyInterval: 5,
		Threshold:      10,
		Timezone:       "Asia/Tokyo",
	},
}

var notifier = faultline.NewNotifier("faultline_go_project", os.Getenv("FAULTLINE_MASTERKEY"), os.Getenv("FAULTLINE_ENDPOINT"), notifications)

func main() {
	defer notifier.Close()
	defer notifier.NotifyOnPanic()
	result := utils.NoticeIcalCalendar()
	if result == nil {
		fmt.Println("done")
	} else {
		fmt.Println(result)
	}
}
