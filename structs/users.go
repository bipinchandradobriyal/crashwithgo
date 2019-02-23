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