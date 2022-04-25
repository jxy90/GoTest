package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_Time(t *testing.T) {
	PublishTime := "2022-04-24T09:36:38.576Z"
	time1, err := CapturedTime2LocalTime(PublishTime)
	if err != nil {
		t.Log(err)
	}
	t.Log(time1)
}

func CapturedTime2LocalTime(captureTime string) (time.Time, error) {
	localTime, err := time.ParseInLocation(time.RFC3339, captureTime, time.Local)
	if err != nil {
		return time.Time{}, fmt.Errorf("parse time error %+v", err)
	}
	return localTime, nil
}
