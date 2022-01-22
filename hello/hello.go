package main

import (
	"fmt"
	"log"

	//"example.com/please"
	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//name := "Max"
	names := []string{"Max", "Paola"}

	//message, err := please.Please(name)
	//message, err := greetings.Hello(name)
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(message)
	fmt.Println(messages)
}
