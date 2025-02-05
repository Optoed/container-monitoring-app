package main

import (
	"github.com/joho/godotenv"
	"log"
	"sync"
	"time"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	var mu sync.Mutex
	var wg sync.WaitGroup

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:

			ips, err := getContainers()
			if err != nil {
				log.Printf("Error getting containers: %v", err)
				continue
			}

			for _, ip := range ips {
				wg.Add(1)

				go func(ip string) {
					defer wg.Done()

					status, pingDuration, err := pingContainer(ip)
					if err != nil {
						log.Printf("Error pinging container %s: %v", ip, err)
						return
					}

					container := &Container{
						IP:           ip,
						Status:       status,
						PingDuration: pingDuration,
						LastPingTime: time.Now(),
					}

					mu.Lock()
					err = sendPingResult(container)
					mu.Unlock()

					if err != nil {
						log.Printf("Error sending ping result for %s: %v", ip, err)
					} else {
						log.Printf("Container %s is %s", ip, status)
					}
				}(ip)
			}

			wg.Wait()
		}
	}
}
