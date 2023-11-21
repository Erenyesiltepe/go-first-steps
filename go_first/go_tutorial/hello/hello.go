package main

import (
	"example/greetings"
	"fmt"
	"log"
)

func main() {
    // Get a greeting message and print it.
    message, err := greetings.Hello("fdsvgd")
    if err!=nil{
        log.Fatal(err)
    }
    fmt.Println(message)
}
