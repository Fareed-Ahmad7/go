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

const (
	ab = iota
	bb
	cb
)

type doctor struct {
	number int
	actorName string
	companions []string

}

type Animal struct {
	
	name string `max:"100"`
	origin string
}

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

	// boolean in go 
	var r bool = true
	fmt.Printf("%v,%T", r, r)  

	// logical tests
	s := 1 == 1
	t := 1 == 2
	fmt.Printf("\n%v,%T\n", s, s)  
	fmt.Printf("%v,%T", t, t) 
	
	// default value of bool is false
	var u bool
	fmt.Printf("%v,%T", u, u)

	// strings
	v := "this is a string"
	w := []byte(v)
	fmt.Printf("%v,%T", w, w)

	// rune
	x := 'a'  // a rune is an integer 32
	fmt.Printf("%v,%T", x, x)  // output { 97, int32 }

	var y rune = 'a'
	fmt.Printf("%v,%T", y, y) 

	// iota   // iota = 0 and next iota is +1
	fmt.Printf("\n%v,%T\n", bb, bb) 

	// arrays
	marks := [3]int{23,57,98}
	marks1 := [...]int{23,57,98,73}
	
	var marks2 [3]string
	marks2[0] = "farid"
	marks2[1] = "farid1"
	marks2[2] = "farid2"

	fmt.Printf("\n%v\n", marks) 
	fmt.Printf("\n%v\n", marks1) 
	fmt.Printf("\n%v\n", marks2) 
	fmt.Printf("\n%v\n", marks2[2]) 
	fmt.Printf("%v\n", len(marks2)) 

	// array matrix
	var array [3][3]int
	array[0] = [3]int{1,2,3}
	array[1] = [3]int{11,22,33}
	array[2] = [3]int{111,222,333}

	fmt.Printf("\n%v\n", array) 
	
	// pointers
	arr := [...]int{1,2,4}
	arr1 := arr
	arr1[1] = 5

	arr2 := [...]int{4,9,6}
	arr3 := &arr2
	arr3[1] = 5

	fmt.Printf("\n%v\n", arr)
	fmt.Printf("\n%v\n", arr1)
	fmt.Printf("\n%v\n", arr3)

	// map
	statePopulations := map[string]int{
		"california" : 329490,
		"los angles" : 498338,
		"new york" : 437820,
		"texas" : 986489,
		"alabama" : 109937,
	}
	fmt.Printf("\n%v\n", statePopulations)
	
	statePopulations["georgia"] = 438239
	fmt.Printf("\n%v\n", statePopulations)
	
	delete(statePopulations, "georgia")
	fmt.Printf("\n%v\n", statePopulations)
	fmt.Printf("\n%v\n", len(statePopulations))
	
	// check if key is in the map or not
	
	// pop, ok := statePopulations["texs"]
	pop, ok := statePopulations["texas"]
	fmt.Println( pop, ok)

	// structs
	aDoctor:= doctor{
		actorName: "farid",
		number: 3,
		companions: []string{
			"ahmed",
			"shaik",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions[0])

	firstanimal := Animal{
		name: "emu",
		origin: "australia",
	}
	fmt.Println(firstanimal)

	// if statements
	if true {
		fmt.Printf("This is true")
	} else {
		fmt.Printf("This is false")
	}

	num := 30
	num1 := 40
	if num > num1 {
		fmt.Printf("num is greater than num1")
	}
	if num < num1 {
		fmt.Printf("num1 is greater than num")
	}
	if num == num1 {
		fmt.Printf("num is equal to num1")
	} 

	// switch statements

	switch 2 {
	case 1 :
		fmt.Printf("one")
	case 2 :
		fmt.Printf("two")
	case 3 :
		fmt.Printf("three")
	default:
		fmt.Printf("none")
	}

	// loops
	for i := 0; i<5; i++{
		fmt.Printf("meow")
	}

	for i,j := 0, 0;  i<5;  i,j = i+1, j+2{
		fmt.Println(i, j)
	} 
	
	fmt.Println("")

	Loop:
			for i := 1 ; i <= 3; i++{
				for j := 1; j <= 3; j++ {
					fmt.Println(i * j)
					if i * j >= 3{
						// break doesn't work as intended here
						break Loop

					}
				}
			}
	
	// looping over collections
	slice := []int{1,2,3}
	fmt.Println(slice)
	for key, value := range slice{
		fmt.Println(key, value)
	}
	for _, value := range slice{
		fmt.Println(value)
	}

	// defer functions
	// defer is executed after completion of main func execution
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")

	// multiple defers are executed in lifo order {last in first out}
	defer fmt.Println("start1")
	defer fmt.Println("middle2")
	defer fmt.Println("end3")

	// defer takes first initialized value of variable
	var1 := "farid"
	defer fmt.Println(var1)
	var1 = "ahmed"
	fmt.Println(var1)

	// panic
	// below code will make go run into exception which makes it no longer execute next lines of code
	// this is called as panic in golang.
	num_1 , num_2 := 1, 0
	ans := num_1/num_2
	fmt.Println(ans)
	// output {
		// panic: runtime error: integer divide by zero
	// }
	
	// using panic message
	fmt.Println("hi")
	panic("something bad happend")
	fmt.Println("hello")

	// after panic is executed defer statements are executed
	
}

