package main

import (
	"log"
	"pinger/models"
	"pinger/service"
	"sync"
	"time"
)

func main() {
	var (
		mu sync.Mutex
		wg sync.WaitGroup
	)

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	log.Println("ticker has been started")

	countOfIterations := 0
	for {
		countOfIterations++
		log.Printf("iteration №%v\n", countOfIterations)
		select {
		case <-ticker.C:
			ips, err := service.GetContainers()
			if err != nil {
				log.Printf("Error getting containers: %v", err)
				continue
			}
			log.Println("IPs of containers = ", ips, " and len(ips) = ", len(ips))

			for _, ip := range ips {
				// log.Printf("№%v ip = %v\n", i, ip)
				wg.Add(1)

				go func(ip string) {
					// log.Println("in go func(ip string) with ip = ", ip)
					defer wg.Done()

					status, pingDuration, err := service.PingContainer(ip)
					log.Println("status, pingDuration, err := service.PingContainer(ip) : ", status, pingDuration, err)
					if err != nil {
						log.Printf("Error pinging container %s: %v", ip, err)
						return
					}

					container := &models.Container{
						IP:           ip,
						Status:       status,
						PingDuration: pingDuration,
						LastPingTime: time.Now(),
					}

					log.Println("created container : ", container.IP, container.Status, container.LastPingTime, container.PingDuration)

					mu.Lock()
					err = service.SendPingResult(container)
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
