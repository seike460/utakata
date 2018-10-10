package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	ical "github.com/lestrrat-go/ical"
	"github.com/nlopes/slack"
)

// SlackSend send Message to Slack
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
				Title: "予定",
				Value: task,
			},
			slack.AttachmentField{
				Title: "時間",
				Value: start,
			},
		},
	}
	params.Attachments = []slack.Attachment{attachment}
	params.Username = "Utakata (goroutine)"
	params.IconEmoji = ":cloud:"
	api.PostMessage(channel, "", params)
}

// getIcalData get icals data
func getIcalData(ical int) io.ReadCloser {
	icalURL := os.Getenv("UTAKATA_ICAL_URLS")
	icalUserName := os.Getenv("UTAKATA_ICAL_USERS")
	icalPass := os.Getenv("UTAKATA_ICAL_PASS")

	if icalURL == "" || icalUserName == "" || icalPass == "" {
		panic("必要な環境変数が設定されていません")
	}

	req, _ := http.NewRequest("GET", icalURL, nil)
	req.Header.Set("Authorization", "Bearer access-token")
	req.SetBasicAuth(icalUserName, icalPass)

	client := new(http.Client)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return resp.Body
}

// NoticeIcalCalendar entrypoint
func NoticeIcalCalendar() error {
	// 仮ループ用
	icals := []int{0}

	wg := &sync.WaitGroup{}

	// goroutine用channel
	icalChan := make(chan io.ReadCloser)

	var icalBody io.ReadCloser

	for ical := range icals {
		go func(icalChan chan io.ReadCloser) {
			icalChan <- getIcalData(ical)
		}(icalChan)
		if icalBody != nil {
			wg.Add(1)
			go func() {
				checkAndSlackSend(icalBody)
				wg.Done()
			}()
		}
		icalBody = <-icalChan
	}
	// 最後の一回分
	if icalBody != nil {
		checkAndSlackSend(icalBody)
	}
	wg.Wait()
	return nil
}

func checkAndSlackSend(icalBody io.ReadCloser) error {
	p := ical.NewParser()
	c, err := p.Parse(icalBody)

	if err != nil {
		return err
	}

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
					return nil
				}
			}
		}
	}
	return nil
}
