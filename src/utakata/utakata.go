package main

import (
	"fmt"

	utils "github.com/seike460/utakata/src"
)

func main() {
	result := utils.GetIcalCalendar()

	fmt.Println(result)

	//for _, item := range result.Items {
	//	layout := "20060102T150405"
	//	sendTime := strings.Replace(*item["dateTime"].S, "-", "", -1)
	//	sendTime = strings.Replace(sendTime, ":", "", -1)
	//	t, err := time.Parse(layout, sendTime)
	//	if err == nil {
	//		now := time.Now().UTC().Add(time.Duration(9) * time.Hour)
	//		addTime := t.Add(time.Duration(5) * time.Minute)
	//		minusTime := t.Add(-time.Duration(5) * time.Minute)
	//		if now.Before(addTime) && now.After(minusTime) {
	//			//utils.SlackSend(*item["name"].S, t.String())
	//		}
	//	}
	//}
}
