package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")

	age :=22   //types are locked at compile time
	fmt.Println(age)
	// age:="sai kiran"  // cant assigin


	 result := add(1, 2)
    fmt.Println(result)
}


func add(a int,b int ) int {
	return a + b
}

