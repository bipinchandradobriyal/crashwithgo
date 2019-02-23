package main

import (
	"fmt"
	"crashwithgo/structs"
)


func main() {

	var _ structs.User

	user := structs.User{Name:"Bipin",Age:32, Address: structs.Address{City: "", State: ""}, DOB:32423434322323}

	// Order needs to maintain
	// Without private field compile time error
	// With private field runtime error "implicit assignment of unexported field 'password'"
	user = structs.User{"Bipin",32, 32423434322323,structs.Address{City: "Delhi", State: "Delhi"},"password"}

	// Assigning address. user1 is a *User type
	user1 := &structs.User{Name:"Bipin",Age:32, Address: structs.Address{City: "NOIDA", State: "UP"}, DOB:32423434322323}

	// Other way to declare
	user2 := new(structs.User)
	*user2 = structs.User{Name:"Bipin",Age:32, Address: structs.Address{City: "", State: ""}, DOB:32423434322323}

	fmt.Println(user,"\n",*user1)
}