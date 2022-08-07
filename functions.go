package main

import (
	"fmt"
)

func main() {
	message()
	name:= "fareed"
	person(name)
	fmt.Println(name)
}

func message(){
	fmt.Println("salaam!")
}
func person(personName string){
	fmt.Println(personName)
	personName = "ahmed"
	fmt.Println(personName)
}