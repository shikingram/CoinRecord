package main

import "time"

const timeTemplate1 = "2006-01-02 15:04:05" //常规类型

func StringToTime(t1 string) time.Time {
	if len(t1) < len(timeTemplate1) {
		t1 += " 00:00:00"
	}
	stamp, _ := time.ParseInLocation(timeTemplate1, t1, time.Local)
	return stamp
}

func TimeFormat(t time.Time) string {
	return t.Format(timeTemplate1)
}

func timeSub(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)

	return int(t1.Sub(t2).Hours() / 24)
}
