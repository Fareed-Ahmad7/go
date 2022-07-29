package main

import (
	"fmt"
	"strconv"
)

// declaring global variables
var m int = 55

// m:= 55  doesn't work here

// simple ways to declare multible variables in go
var (
	name string = "fareed"
	class int = 13
	gpa float32 = 7.5
)


func main() {
	fmt.Println("salaam!")
	fmt.Println(349)
	var i int
	i = 100
	fmt.Println(i)
	i = 7
	fmt.Println(i)

	// simple ways to initialize var
	var j int = 8
	fmt.Println(j)

	// using : so that go can figure-out the datatype itself
	k := 9
	fmt.Println(k)
	fmt.Printf("%v, %T,", k, k)

	var l float32 = 9
	fmt.Printf("%v, %T", l, l)

	// access global var
	fmt.Println(m)
	fmt.Println(name)
	fmt.Println(class)

	// type conversion
	var n int = 22
	fmt.Printf("%v,%T", n, n)

	var o float32
	o = float32(n)
	fmt.Printf("%v,%T", o, o)

	// string conversion
	// var p string
	// p = string(n)
	// fmt.Println(" ")
	// fmt.Printf("%v,%T", p, p)  
	// output : *, string

	// using strconv to convert int to ascii
	var q string
	q = strconv.Itoa(n)
	fmt.Printf("%v,%T", q, q)  
}