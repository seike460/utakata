package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	ical "github.com/lestrrat-go/ical"
	"github.com/nlopes/slack"
)

func SlackSend(task string, start string) {

	token := os.Getenv("UTAKATA_SLACK_TOKEN")
	channel := os.Getenv("UTAKATA_SLACK_CHANNEL")

	if token == "" || channel == "" {
		panic("必要な環境変数が設定されていません")
	}

	api := slack.New(token)
	params := slack.PostMessageParameters{}
	attachment := slack.Attachment{
		Fields: []slack.AttachmentField{
			slack.AttachmentField{
				Title: "タスク",
				Value: task,
			},
			slack.AttachmentField{
				Title: "時間",
				Value: start,
			},
		},
	}
	params.Attachments = []slack.Attachment{attachment}
	params.Username = "Utakata"
	params.IconEmoji = ":alarm_clock:"
	channelID, timestamp, err := api.PostMessage(channel, "", params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}

/**
 *
 */
func NoticeIcalCalendar() io.ReadCloser {

	icalUrl := os.Getenv("UTAKATA_ICAL_URLS")
	icalUserName := os.Getenv("UTAKATA_ICAL_USERS")
	icalPass := os.Getenv("UTAKATA_ICAL_PASS")

	if icalUrl == "" || icalUserName == "" || icalPass == "" {
		panic("必要な環境変数が設定されていません")
	}

	req, _ := http.NewRequest("GET", icalUrl, nil)
	req.Header.Set("Authorization", "Bearer access-token")
	req.SetBasicAuth(icalUserName, icalPass)

	client := new(http.Client)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	p := ical.NewParser()
	c, err := p.Parse(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	// snip
	for e := range c.Entries() {
		ev, ok := e.(*ical.Event)
		if !ok {
			continue
		}
		summary, ret := ev.GetProperty("summary")
		if ret == true {
			if len(summary.Parameters()["VALUE"]) > 0 {
				fmt.Println(summary.Parameters())
			}
		}
		dtstart, ret := ev.GetProperty("dtstart")
		if len(dtstart.RawValue()) > 10 {
			layout := "20060102T150405"
			t, err := time.Parse(layout, dtstart.RawValue())
			if err == nil {
				now := time.Now().UTC().Add(time.Duration(9) * time.Hour)
				addTime := t.Add(time.Duration(5) * time.Minute)
				minusTime := t.Add(-time.Duration(5) * time.Minute)
				if now.Before(addTime) && now.After(minusTime) {
					SlackSend(summary.RawValue(), t.String())
				}
			}
		}
	}
	return resp.Body
}
