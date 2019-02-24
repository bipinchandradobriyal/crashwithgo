package main

import (
	"fmt"
	"crashwithgo/structs"
)




type BBRequest interface {
	PrintData()
	//RequestHandler()
}

func checkTheValueByInterface(value BBRequest) {
	value.PrintData()
}

func main() {
	fmt.Println("\n\n\n*******************************")
	req := structs.Request1{"John",12}
	checkTheValueByInterface(&req)

	fmt.Println("\n\n\n*******************************")
	req2 := structs.Request2{"Tom",21,"Noida"}
	checkTheValueByInterface(&req2)

}
