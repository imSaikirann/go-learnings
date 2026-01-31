package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")

	age :=22   //types are locked at compile time
	fmt.Println(age)
	// age:="sai kiran"  // cant assigin


	 result := add(1, 2)
    fmt.Println(result)
	
	user := User{    // assigiing data 
        Name: "Kiran",
		Age: 22,
	}

	fmt.Println(user)
}


func add(a int,b int ) int {
	return a + b
}


type User struct {   // a blue print 
	Name string
	Age  int
}

