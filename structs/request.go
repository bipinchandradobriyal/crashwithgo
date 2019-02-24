package structs

import "fmt"

type Request1 struct {
	Name 	string
	Age		int
}
func (r *Request1) PrintData(){
	fmt.Println(r.Name , r.Age)
}
func (r Request1) RequestHandler(){
}


type Request2 struct {
	Name 	string
	Age		int
	City	string
}
func (r *Request2) PrintData(){
	fmt.Println(r.Name ,r.Age, r.City)
}

