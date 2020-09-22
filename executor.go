package main

import (
	"os"
	"time"
)

func sayHello() string {
	return ("Hello from your Docker build agent!!")
}

func getTime() string {
	currentTime := time.Now()
	return currentTime.Format("2006.01.02 15:04:05")
}

func getHostName() string {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return name
}
