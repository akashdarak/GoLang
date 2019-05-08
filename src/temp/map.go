package main

import (
	"fmt"
)

func main() {
	jb := []string{"James", "Bond", "Movies"}
	//fmt.Println(jb)
	mp := []string{"Ethan", "Hunt", "Pictures"}
	//fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)

	//Maps store unordered key-value pairs
	m := map[string]int{
		"Hello": 4,
		"World": 5,
	}
	fmt.Println(m)
	fmt.Println(m["Hello"])
	fmt.Println(m["No_key"]) //Prints value 0 of there is no key stored
}

