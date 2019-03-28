package main

import "fmt"

const(
	year = 2019 + iota
	yearn = 2019 + iota
	a = 42				//Untyped constant
	b int = 43			//Typed constant
)

func main(){
	fmt.Println(year, yearn)
	n := 25
	fmt.Println(n)
	fmt.Printf("\t%b", n)		//Binary
	fmt.Printf("\t%#x\n", n)	//Hex

	m := n << 1					//Bit shift
	fmt.Printf("\t%b", m)

	t := (42 == 42)				//Assign expression values
	fmt.Println("\t", t)

	fmt.Println("\n", a, b)
}