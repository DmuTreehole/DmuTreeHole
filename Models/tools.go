package models

import (
	"time"
)

func GetDatetime() (datetime string) {
	templage := "2006-01-02 15:04:05"
	datetime = time.Now().Format(templage)
	return
}
