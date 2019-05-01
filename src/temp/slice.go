package main

import (
	"fmt"
)

func main() {
	//x := type{values}	//Composite Literal
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(len(x))
	fmt.Println(x)
	fmt.Println(x[2:])	//Slicing a slice
	fmt.Println(x[0])
	
	for i, v := range x{
		fmt.Println(i, v)
	}
	
	x = append(x, 77, 88)
	fmt.Println(x)
	
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}

//Slice allows values of same type
