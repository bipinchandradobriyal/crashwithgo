package structs

import "fmt"

type Company struct{
	name 		string
	address		string
}

func (self *Company) SetName(name string){
	fmt.Println(self)
	self.name = name
}

func (self Company) SetAddress(address string){
	fmt.Println(self)
	self.address = address
}

func (self *Company) GetName() (string){
	fmt.Println(self)
	return self.name
}

func (self Company) GetAddress() (string){
	fmt.Println(self)
	return self.address
}

func (self *Company) ToString() (string){
	fmt.Println("Inside ToString")
	return self.name + " " + self.address
}

