package main

import (
	"bytes"
	"encoding/json"
	"github.com/go-ping/ping"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

func getContainers() ([]string, error) {
	cmd := exec.Command("docker", "ps", "-q")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	containersID := string(output)
	ids := strings.Fields(containersID)

	var ips []string
	for _, id := range ids {
		psCmd := exec.Command("docker", "inspect", "--format", "{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}", id)
		ip, err := psCmd.Output()
		if err != nil {
			log.Printf("Error getting IP for container %s: %v", id, err)
			continue
		}
		ips = append(ips, strings.TrimSpace(string(ip)))
	}

	return ips, nil
}

func pingContainer(ip string) (string, time.Duration, error) {
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
	if stats.PacketLoss > 0 {
		return "down", stats.AvgRtt, nil
	}

	return "alive", stats.AvgRtt, nil
}

func sendPingResult(container *Container) error {
	url := os.Getenv("BACKEND_URL") + "/containers" // URL backend API
	data, err := json.Marshal(container)
	if err != nil {
		return err
	}
	_, err = http.Post(url, "application/json", bytes.NewBuffer(data))
	return err
}
