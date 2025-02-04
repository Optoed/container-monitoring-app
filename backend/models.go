package main

type Container struct {
	ID           int    `json:"id"`
	IP           string `json:"ip"`
	Status       string `json:"status"`
	LastPingTime string `json:"last_ping_time"`
}
