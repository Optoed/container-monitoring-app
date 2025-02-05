package models

import "time"

type Container struct {
	ID           int           `json:"id" db:"id"`
	IP           string        `json:"ip" db:"ip"`
	Status       string        `json:"status" db:"status"`
	LastPingTime time.Time     `json:"last_ping_time" db:"last_ping_time"`
	PingDuration time.Duration `json:"ping_duration" db:"ping_duration"`
}
