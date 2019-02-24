package main

import (
	"fmt"
	"reflect"
	"crashwithgo/structs"
)


func checkTheValue(value interface{}) {

	fmt.Println("Value:: ",value)

	switch value.(type) {
	case structs.Request1:
		fmt.Println("****This is request1 type")
		val := value.(structs.Request1)
		(&val).PrintData()

	case structs.Request2:
		fmt.Println("****This is request2 type")
		val := value.(structs.Request2)
		(&val).PrintData()

	default:
		fmt.Println("Wrong argument")
	}

}


// Identifying the type of the argument using reflection
func Any(value interface{}) string {
	return checkValueByReflection(reflect.ValueOf(value))
}

func checkValueByReflection(value reflect.Value) string{

	switch value.Kind(){
	case reflect.Ptr:
		fmt.Println("Pointer value")
	case reflect.Struct:
		fmt.Println("Struct value")
	}

	return "Not able to indentify type"
}

func main() {
	fmt.Println("\n\n\n*******************************")
	req := structs.Request1{"John",32}
	checkTheValue(req)
	checkTheValue(nil)

	fmt.Println("\n\n\n*******************************")
	req2 := structs.Request2{"Ram",32,"Noida"}
	checkTheValue(req2)
	Any(req2)

}


