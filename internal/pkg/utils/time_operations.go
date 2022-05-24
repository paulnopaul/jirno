package utils

import "time"

func GetDayRange(t time.Time) (time.Time, time.Time) {
	start := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	end := start.AddDate(0, 0, 1).Add(time.Second * -1)
	return start, end
}
