package models

import "time"

type URL struct {
	Name string
	Date time.Time
	Tags []string
	Link string
}
