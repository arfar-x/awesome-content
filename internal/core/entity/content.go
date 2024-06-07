package entity

import "time"

type Content struct {
	Name      string
	Title     string
	Rate      int
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
