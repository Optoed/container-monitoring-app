package service

import (
	"github.com/go-ping/ping"
	"log"
	"os/exec"
	"strings"
	"time"
)

func GetContainers() ([]string, error) {
	cmd := exec.Command("docker", "ps", "-aq")
	output, err := cmd.Output()
	// log.Printf("output, err := cmd.Output(); output = %s, err = %v\n", output, err)
	if err != nil {
		return nil, err
	}

	containersID := string(output)
	// log.Println("containersID := string(output); containersID = ", containersID)
	ids := strings.Fields(containersID)
	log.Println("IDs of containers = ", ids)

	var ips []string
	for _, id := range ids {
		//log.Println("id in range ids = ", id)
		psCmd := exec.Command("docker", "inspect", "--format", "{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}", id)
		ip, err := psCmd.Output()
		//log.Println("ip = ", string(ip))
		if err != nil {
			log.Printf("Error getting IP for container %s: %v", id, err)
			continue
		}
		ips = append(ips, strings.TrimSpace(string(ip)))
	}

	log.Println("IPs of containers = ", ips)
	return ips, nil
}

func PingContainer(ip string) (string, time.Duration, error) {
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		return "", 0, err
	}

	pinger.Count = 1                 // Будет отправлен 1 пакет
	pinger.Timeout = time.Second * 2 // Если пинг не ответит в течение 2 секунд, запрос считается неудачным.

	err = pinger.Run()
	if err != nil {
		return "down", 0, err
	}

	stats := pinger.Statistics()
	pingDuration := stats.AvgRtt

	if stats.PacketLoss > 0 {
		return "down", pingDuration, nil
	}

	return "alive", pingDuration, nil
}
