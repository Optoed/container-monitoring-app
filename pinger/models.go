package main

import "time"

type Container struct {
	ID           int           `json:"id"`
	IP           string        `json:"ip"`
	Status       string        `json:"status"`
	LastPingTime time.Time     `json:"last_ping_time"`
	PingDuration time.Duration `json:"ping_duration"`
}
