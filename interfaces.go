package main

import (
	"fmt"
)

func main(){
	fmt.Println("interfaces")
}

type Writer interface {
	Write([]byte) (int, error)
}