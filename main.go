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

func main() {
	status := CheckStatus()
	fmt.Println(time.Now(), status.String())
	for {
		time.Sleep(time.Second)
		if newStatus := CheckStatus(); newStatus != status {
			fmt.Println(time.Now(), newStatus.String())
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
