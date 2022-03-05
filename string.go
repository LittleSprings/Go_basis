package main

import (
	"fmt"
	"strings"
)

func func_string() {
	var str string = "Hi, I'm Marc, Hi."

	fmt.Printf("The position of \"Marc\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Marc"))

	fmt.Printf("The position of the first instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Hi"))
	fmt.Printf("The position of the last instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))

	fmt.Printf("The position of \"Burger\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Burger"))

	var str1 string = "Hello, how is it going, Hugo?"
	// var str2 string

	fmt.Printf("Number of H's in %s is: ", str1)
	fmt.Printf("%d\n", strings.Count(str1, "H"))

	fmt.Printf("%s\n", strings.Replace(str1, "oo", "e", 2))
}
