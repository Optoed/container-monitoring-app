package main

import (
	"log"
	"pinger/models"
	"pinger/service"
	"sync"
	"time"
)

func main() {
	//if err := godotenv.Load(); err != nil {
	//	log.Fatal("Error loading .env file")
	//}

	var mu sync.Mutex
	var wg sync.WaitGroup

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	log.Println("ticker has been started")

	countOfInterations := 0
	for {
		countOfInterations++
		log.Printf("iteration №%v\n", countOfInterations)
		select {
		case <-ticker.C:
			log.Println("before getting ips, err := service.GetContainers()")

			ips, err := service.GetContainers()
			if err != nil {
				log.Printf("Error getting containers: %v", err)
				continue
			}

			log.Println("ips:", ips, len(ips))

			for i, ip := range ips {

				log.Printf("№%v ip = %v\n", i, ip)

				wg.Add(1)

				go func(ip string) {
					log.Println("in go func(ip string) with ip = ", ip)

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

					log.Println("created container : ", container)

					mu.Lock()
					log.Println("before err = service.SendPingResult(container)")
					err = service.SendPingResult(container)
					log.Println("after err = service.SendPingResult(container); err = ", err)
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
