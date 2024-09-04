package main

import "fmt"

func printName() {
	name := "Paras"
	fmt.Println(name)
}

func calculateSum(){
	var a int = 20
	var b  int = 30
	fmt.Println(a+b)
}

// arrays in golang

var arr = [4]int{1,2,3,4}

// main function which executes other code

func main() {
	fmt.Println(arr)
	// age := 20
	// fmt.Println("Age:", age)
	// fmt.Println("Hello, World!")
	// printName()
	// calculateSum()
}
