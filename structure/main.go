package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// vaishnavi := person{
	// 	firstName: "Vaishnavi",
	// 	lastName:  "Disale",
	// }

	// vaishnavi := person{"Vaishnavi","Disale"}

	var vaishnavi person

	fmt.Println(vaishnavi)
	fmt.Printf("%+v", vaishnavi)

}