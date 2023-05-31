package main

import (
	"fmt"
	"net"
	"net/http"
	"os/exec"
	"regexp"
	"strings"
	"time"
)


func testICMP() (status string, msg string, pingTime string) {
	out, err := exec.Command("ping", "-c", "3", *targetUrl).Output()
	if err != nil {
		return "down", fmt.Sprintf("Error: %v", err), "unknown"
	}

	output := string(out)
	status = "down"
	if strings.Contains(output, "received") {
		status = "up"
	}

	packetLossRe := regexp.MustCompile(`(\d+)% packet loss`)
	packetLossMatch := packetLossRe.FindStringSubmatch(output)
	packetLoss := "unknown"
	if len(packetLossMatch) > 0 {
		packetLoss = packetLossMatch[1]
	}

	rttRe := regexp.MustCompile(`rtt min/avg/max/mdev = [\d\.]+/([\d\.]+)/[\d\.]+/[\d\.]+ ms`)
	rttMatch := rttRe.FindStringSubmatch(output)
	pingTime = "unknown"
	if len(rttMatch) > 0 {
		pingTime = rttMatch[1]
	}

	msg = fmt.Sprintf("Packet loss: %s%%", packetLoss)

	return status, msg, pingTime
}
 
func testHTTP(protocol string) (status string, msg string, pingTime string) {
	startTime := time.Now()

	resp, err := http.Get(fmt.Sprintf("%s://%s", protocol, *targetUrl))
	if err != nil {
		return "down", fmt.Sprintf("Error: %v", err), "unknown"
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "down", fmt.Sprintf("Status code: %d", resp.StatusCode), "unknown"
	}

	pingDuration := time.Since(startTime)

	return "up", "OK", fmt.Sprintf("%.2f ms", float64(pingDuration.Microseconds())/1000.0)
}


func testTCP() (status string, msg string, pingTime string) {
	startTime := time.Now()

	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", *targetUrl, *port), time.Second*3)
	if err != nil {
		return "down", fmt.Sprintf("Error: %v", err), "unknown"
	}
	defer conn.Close()

	pingDuration := time.Since(startTime)

	return "up", "OK", fmt.Sprintf("%.2f ms", float64(pingDuration.Microseconds())/1000.0)
}
