package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)


func main(){

	rand.Seed(time.Now().UnixNano())

	target := rand.Intn(100) +1
	maxAttempts :=7
	attempts :=0

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("🎮 Welcome to the Guessing Game!")
	fmt.Println("I have picked a number between 1 and 100.")
	fmt.Println("You have", maxAttempts, "attempts.")
	fmt.Println("----------------------------------")

	for attempts < maxAttempts {
		fmt.Print("Enter your guess")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		guesss , err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("❌ Please enter a valid number.")
			continue
		}

		attempts++

		if guesss < target {
			fmt.Println("low")
		} else if guesss > target {
			fmt.Println("high")
		} else {
			fmt.Printf("correct %d",attempts)
			return
		}

	}
}