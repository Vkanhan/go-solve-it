package main

import (
	"fmt"
	"strings"
)

func countLogs(logs []string) (int, int) {
	var (
		infoCount, errCount int
	)
	for _, log := range logs {
		var (
			pipeCount int
			level     = ""
		)
		for i := 0; i < len(log); i++ {
			if log[i] == '|' {
				pipeCount++
				continue
			}

			if pipeCount == 1 && log[i] == ' ' {
				continue
			}

			if pipeCount == 1 {
				for i < len(log) && log[i] != '|' {
					level += string(log[i])
					i++
				}
				break
			}
		}
		switch strings.TrimSpace(level) {
		case "INFO":
			infoCount++
		case "ERROR":
			errCount++
		}
	}
	return infoCount, errCount
}

func count(logs []string) (int, int) {
	var (
		infoCount, errCount int 
	)
	for _, log := range logs {
		if strings.Contains(log, "| INFO |") {
			infoCount++
		}

		if strings.Contains(log, "| ERROR |") {
			errCount++
		}
	}
	return infoCount, errCount
}

func countlogs(logs []string) (int, int) {
	var (
		infoCount, errCount int 
	)

	for _, log := range logs {
		parts := strings.Split(log, "|")
		if len(parts) > 2 {
			level := strings.TrimSpace(parts[1])
			switch level {
			case "INFO":
				infoCount++
			case "ERROR":
				errCount++
			}
		}
	}
	return infoCount, errCount 
}

func main() {
	logs := []string{
		"[2024-01-19 12:30:45] | INFO | User: 123 | Action: visited | resource: page/home",
		"[2024-01-19 12:30:45] | INFO | User: 123 | Action: exit | resource: page/home",
		"[2024-01-19 12:30:45] | ERROR | User: 234 | Action: visited | resource:page/analytics | Error details: Page not found: /404",
		"[2024-01-19 12:30:45] | WARN | User: 456 | Action: update | resource:page/home | Warn details: unauthorized change",
	}

	fmt.Println(countLogs(logs))
	fmt.Println(count(logs))
	fmt.Println(countlogs(logs))
}
