package main

import "fmt"

func main(){
	for n:=10; n<=100; n++{
		fmt.Println(n%4)
	}

	for i:=65; i<=90; i++{
		fmt.Printf("%#U\n", i)
	}
}