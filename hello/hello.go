package main 

import (
	"fmt"
	"go-server/greetings"
)


func main() {
	message:= greetings.Hello("Sai Kiran")
	fmt.Println(message)
}