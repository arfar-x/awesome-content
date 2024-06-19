// Package entity provides a general domain entity for other entities passed across
// application layers.
package entity

import "time"

type Entity struct {
	ID        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}
