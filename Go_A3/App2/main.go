package main

import "fmt"

func main() {
	name := "Pranay"
	message := "Hello"
	age := 24
	greeting(message, name, age)
}
func greeting(message string, name string, age int) {
	fmt.Println(message, name, ". You are ", age, "years old")
}
