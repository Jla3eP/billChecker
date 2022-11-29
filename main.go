package main

import (
	"fmt"
	"net/http"
	"time"
)

type Status int

func (s Status) String() string {
	if s == Enable {
		return "Enable"
	}
	return "Disable"
}

const (
	Enable = iota
	Disable
)

var client = http.Client{
	Timeout: 1 * time.Second,
}

func CheckStatus() Status {
	_, err := client.Get("http://www.univ.kiev.ua/")
	if err != nil {
		return Disable
	}
	return Enable
}

func PrintTimeAndStatus(status Status) {
	tn := time.Now()
	t := time.Date(tn.Year(), tn.Month(), tn.Day(), tn.Hour(), tn.Minute(), tn.Second(), 0, time.FixedZone("UTC+2", 2*60*60))
	text := []byte("")
	text = t.AppendFormat(text, "MST 02.01.2006 15:04:05")
	fmt.Println(string(text), status)
}

func main() {
	status := CheckStatus()
	PrintTimeAndStatus(status)
	for {
		time.Sleep(time.Second)
		if newStatus := CheckStatus(); newStatus != status {
			PrintTimeAndStatus(newStatus)
			dur := 5 * time.Second
			deadLine := time.Now().Add(dur)
			for deadLine.After(time.Now()) {
				time.Sleep(250 * time.Millisecond)
				fmt.Print("\a")
			}
			status = newStatus
		}
	}
}
