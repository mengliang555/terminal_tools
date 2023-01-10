package data_info

import "time"

type TimeTickerStruct struct {
	Key      string
	Duration time.Time
	Behave   func()
}
