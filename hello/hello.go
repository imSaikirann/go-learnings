package main 

import (
	"fmt"
	"go-server/greetings"
	"log"
	
)


func main() {
	message, err := greetings.Hello("")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}