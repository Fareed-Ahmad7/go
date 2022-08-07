package main

import( 
	"fmt"
)

func main(){
		a := 33
		b := a
		fmt.Println(a,b)

		var c int = 33		
		var d *int = &c   
		fmt.Println(c,d) // prints c = 33 and d = 0xc000016098
		fmt.Println(c, *d) // prints c = 33 and d = 33
		// change c value dosnt change d bcz d is just a pointer
		c = 44
		fmt.Println(c,*d)
		
		// also if we change the value of d the value of c changes too
		*d = 55
		fmt.Println(c, *d)
}