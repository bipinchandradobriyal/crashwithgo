package structs

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
	Password	string
}

// Struct embedding
type Person struct{
	Address	// This can be pointer
	Name 		string
	Age			int
	DOB			int64
	Password	string
}
