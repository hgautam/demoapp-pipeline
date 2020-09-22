package main

import (
	"os"
	"testing"
)

func TestSayHello(t *testing.T) {
	result := sayHello()
	expectedResult := "Hello from your Docker build agent!!"
	if result != expectedResult {
		t.Errorf("sayHello() failed, expected %v, got %v", expectedResult, result)
	}
}

func TestGetHostName(t *testing.T) {
	result := getHostName()
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	if result != name {
		t.Errorf("getHostName() failed, expected %v, got %v", name, result)
	}
}
