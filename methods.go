package main

import (
	"fmt"
	"crashwithgo/structs"
)

type Employee struct{
	Name 	string
	Age		int
}

func (self Employee) getInfo() (info string, e error){
	info = self.Name + " " + string(self.Age)
	return info, e
}

func (self *Employee) getPtrInfo() (info string, e error){
	info = self.Name + " " + string((*self).Age)
	return info, e
}

func main(){
	emp := Employee{Name:"Tom",Age:12}
	fmt.Println(emp.getInfo())
	fmt.Println(emp.getPtrInfo())


	//
	company := structs.Company{}
	company.SetName("My Company Name")
	company.SetAddress("Company address")
	fmt.Println(company.ToString())

}