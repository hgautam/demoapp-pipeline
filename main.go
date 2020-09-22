package main

import "fmt"
import "time"
import "os"


func main() {
    fmt.Println("Hello from your Docker build agent!!")
    currentTime := time.Now()
    fmt.Println("Current time and date: ", currentTime.Format("2006.01.02 15:04:05"))
    name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
    fmt.Println("Container hostname:", name)
}
