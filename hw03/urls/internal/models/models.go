package models

import "time"

type URL struct {
	Description string
	Date        time.Time
	Tags        []string
	Link        string
}
