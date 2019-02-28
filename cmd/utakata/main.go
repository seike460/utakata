package main

import (
	"log"
	"path/filepath"

	"github.com/faultline/faultline-go/faultline"
	"github.com/seike460/utakata/utakata"
	"github.com/spf13/viper"
)

var notifications = []interface{}{
	faultline.Slack{
		Type:           "slack",
		Endpoint:       viper.GetString("FAULTLINE_NOTIFY_SLACK_ENDPOINT"),
		Channel:        viper.GetString("FAULTLINE_NOTIFY_SLACK_CHANNEL"),
		Username:       "faultline-notify",
		NotifyInterval: 5,
		Threshold:      10,
		Timezone:       "Asia/Tokyo",
	},
}

var notifier = faultline.NewNotifier(
	viper.GetString("FAULTLINE_PROJECT_NAME"),
	viper.GetString("FAULTLINE_MASTERKEY"),
	viper.GetString("FAULTLINE_ENDPOINT"),
	notifications,
)

func main() {
	defer notifier.Close()
	defer notifier.NotifyOnPanic()
	path := filepath.Join(
		"$GOPATH",
		"src",
		"github.com",
		"seike460",
		"utakata")
	viper.AddConfigPath(path)
	viper.SetConfigName("utakata")
	viper.ReadInConfig()

	err := utakata.NoticeIcalCalendar()
	if err != nil {
		log.Fatal(err)
	}
}
