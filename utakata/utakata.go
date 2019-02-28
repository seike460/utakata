package utakata

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/lestrrat-go/ical"
	"github.com/nlopes/slack"
	"github.com/spf13/viper"
)

// SlackType Daliy Hour Minute
var SlackType string

// SlackSend send Message to Slack
func SlackSend(task string, start string) error {

	token := getConfigValue("UTAKATA_SLACK_TOKEN")
	channel := getConfigValue("UTAKATA_SLACK_CHANNEL")
	if token == "" || channel == "" {
		return errors.New("plz set UTAKATA_SLACK_TOKEN & UTAKATA_SLACK_CHANNEL")
	}
	api := slack.New(token)
	params := slack.PostMessageParameters{}
	attachment := slack.Attachment{
		Fields: []slack.AttachmentField{
			slack.AttachmentField{
				Title: "Schedule",
				Value: task,
			},
			slack.AttachmentField{
				Title: "Time",
				Value: start,
			},
		},
	}
	params.Attachments = []slack.Attachment{attachment}
	params.Username = "Utakata"
	params.IconEmoji = ":cloud:"
	_, _, err := api.PostMessage(channel, "", params)
	if err != nil {
		return err
	}
	return nil
}

// getIcalData get icals data
func getIcalData(ical int) io.ReadCloser {
	icalURL := getConfigValue("UTAKATA_ICAL_URLS_" + strconv.Itoa(ical))
	icalUserName := getConfigValue("UTAKATA_ICAL_USERS_" + strconv.Itoa(ical))
	icalPass := getConfigValue("UTAKATA_ICAL_PASS_" + strconv.Itoa(ical))

	if icalURL == "" {
		log.Fatal("plz set UTAKATA_ICAL_URLS_" + strconv.Itoa(ical))
		return nil
	}

	req, err := http.NewRequest("GET", icalURL, nil)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	// Set Basic Auth
	if icalUserName != "" && icalPass != "" {
		req.Header.Set("Authorization", "Bearer access-token")
		req.SetBasicAuth(icalUserName, icalPass)
	}

	client := new(http.Client)
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
		return nil
	}
	return resp.Body
}

// NoticeIcalCalendar entrypoint
func NoticeIcalCalendar() error {

	if getConfigValue("UTAKATA_ICAL_NUM") == "" {
		return nil
	}

	calNum, err := strconv.Atoi(getConfigValue("UTAKATA_ICAL_NUM"))
	if err != nil {
		return err
	}

	if len(os.Args) > 1 && os.Args[1] == "Daily" {
		SlackType = "Daily"
	}

	wg := &sync.WaitGroup{}
	iCalChan := make(chan io.ReadCloser)
	var iCalBody io.ReadCloser

	num := 1
	for {
		go func(iCalChan chan io.ReadCloser) {
			iCalChan <- getIcalData(num)
		}(iCalChan)
		if iCalBody != nil {
			wg.Add(1)
			go func() {
				defer wg.Done()
				err := checkAndSlackSend(iCalBody)
				if err != nil {
					log.Println(err)
				}
			}()
		}
		iCalBody = <-iCalChan
		num++
		if num > calNum {
			break
		}
	}
	// lastExec Once
	if iCalBody != nil {
		err := checkAndSlackSend(iCalBody)
		if err != nil {
			return err
		}
	}
	wg.Wait()
	return nil
}

func checkAndSlackSend(iCalBody io.ReadCloser) error {
	p := ical.NewParser()
	c, err := p.Parse(iCalBody)

	if err != nil {
		return err
	}

	for e := range c.Entries() {
		ev, ok := e.(*ical.Event)
		if !ok {
			continue
		}
		summary, ret := ev.GetProperty("summary")
		if ret != true {
			return errors.New("fail get summary")
		}
		if ret != true {
			return errors.New("fail get dtstart")
		}
		dtstart, ret := ev.GetProperty("dtstart")
		// Non Date

		layout := "20060102T150405"
		if len(dtstart.RawValue()) == 8 {
			if SlackType != "Daily" {
				continue
			}
			layout = "20060102"
		}
		t, err := time.Parse(layout, dtstart.RawValue())
		if err != nil {
			return err
		}
		// @TODO setting from config TimeZone
		now := time.Now().UTC().Add(time.Duration(9) * time.Hour)
		if SlackType == "Daily" {
			if t.Format(layout) == now.Format(layout) {
				err = SlackSend(summary.RawValue(), t.String())
			}
		} else {
			// @TODO setting from config Minute
			addTime := t.Add(time.Duration(5) * time.Minute)
			minusTime := t.Add(-time.Duration(5) * time.Minute)
			if now.Before(addTime) && now.After(minusTime) {
				err = SlackSend(summary.RawValue(), t.String())
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func getConfigValue(configString string) string {
	// serverless for Production
	val := os.Getenv(configString)
	if val != "" {
		return val
	}
	// local for dev
	return viper.GetString(configString)
}
