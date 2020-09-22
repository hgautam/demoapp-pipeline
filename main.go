package main

import "fmt"

func main() {
	fmt.Println(sayHello())
	fmt.Println("Current time and date", getTime())
	fmt.Println("Container host name is: ", getHostName())
}
