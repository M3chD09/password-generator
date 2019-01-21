package main

import (
	. "./core"
	"fmt"
)

func main() {
	fmt.Println("Please enter lower case letters:")
	var text string
	fmt.Scanln(&text)
	if IslowCase(text) {
		output := Generator(text)
		fmt.Println("Generated password:")
		fmt.Println(output)
	} else {
		fmt.Println("Not lower case letters!")
	}
}
