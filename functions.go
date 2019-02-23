package main

import (
	"crashwithgo/structs"
	"errors"
	"fmt"
)

/*
	func func_name(parameter_list) (return_type_lists){
		function body
	}
*/


func passByVal(a int, y structs.User) (n int, e error){

	if a == 0 && y == (structs.User{}) {
		return 1200, errors.New("Insufficient data")
	}
	a++
	fmt.Println("a: ",a, "y: ",y)
	return n,e
}

func passByRef(a *int, y *structs.User) (n int, e error){

	if *a == 0 && *y == (structs.User{}) {
		return 1200, errors.New("Insufficient data")
	}
	*a++
	fmt.Println("a: ",a, "y: ",y)
	return n,e
}


func main() {

	// 1: Pass by value
	var num int
	user := structs.User{Name:"Bipin",Age:32, Address: structs.Address{City: "NOIDA", State: "UP"}, DOB:32423434322323}
	num = 34
	passByVal(num,user)
	fmt.Println("Pass by val num: ",num)
	fmt.Println("**************************************************\n\n\n")


	// 2: Pass by refrence
	num = 34
	passByRef(&num,&user)
	fmt.Println("Pass by ref num: ",num)
	fmt.Println("**************************************************\n\n\n")


	// 3: Anonymous function
	func () {
		// Write some code
	}()


	// 4: Receiving multiple result types
	_,error := passByVal(12,user)
	fmt.Println(error)
	fmt.Println("**************************************************\n\n\n")


	// 5: Anonymous functions
	func() string{return "success from myAnonymousFunc"}()
	myAnonymousFunc := func() string{return "success from myAnonymousFunc"}
	fmt.Println(myAnonymousFunc())
	fmt.Println("**************************************************\n\n\n")


	// 6: Passing a function as argument also known as higher-order functions
	aFunc := func() {fmt.Println("aFunc called")}
	func (x func()){
		x()
	}(aFunc)
	fmt.Println("**************************************************\n\n\n")


	// 7: By deferring any function call, the function will be executed at the end of the execution
	defer func() {fmt.Println("This is deferred function")}()



	// 8: Function panic: if any function panics then the execution stops at that point, all the deferred
	// functions are executed
	// The panic sequence continues all the way up the call stack until the main function is reached and
	// the program exits (crashes).
	panic("Something went wrong")


	// Recovering from panic. The deferred function calls should be at the top of function body
	defer func(){
		if r := recover(); r != nil{
			fmt.Println("Panic recovered:  ",r)
		}
	}()

}

