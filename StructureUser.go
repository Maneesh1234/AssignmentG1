package main

import "fmt"

// Defining a struct type
type User struct {
	Name    string
	city    string
	Pincode int
}

func main() {

	// Declaring a variable of a `struct` type
	// All the struct fields are initialized
	// with their zero value
	var a User
	fmt.Println(a)

	// Declaring and initializing a
	// struct using a struct literal
	a1 := User{"Maneesh", "Lucknow", 225201}

	fmt.Println("User1: ", a1)

	// Naming fields while
	// initializing a struct
	a2 := User{Name: "John", city: "England",
		Pincode: 547555}

	fmt.Println("User2: ", a2)

	// Uninitialized fields are set to
	// their corresponding zero-value
	a3 := User{Name: "Tom"}
	fmt.Println("User3: ", a3)
}
