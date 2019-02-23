package main

import "fmt"

type Address struct{
	City		string
	State		string
	Country		string
	HouseNo		string
	LandMark	string
}

type User struct{
	Name 		string
	Age			int
	DOB			int64
	Address		Address
}


func main()  {

	var _ User

	user := User{Name:"Bipin",Age:32, Address: Address{City: "", State: ""}, DOB:32423434322323}
	user1 := &User{Name:"Bipin",Age:32, Address: Address{City: "", State: ""}, DOB:32423434322323}

	fmt.Println(user,user1)
}