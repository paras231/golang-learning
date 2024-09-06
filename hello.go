package main

import "fmt"
import (
	"os"
	"bufio"
)

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

// func main() {
// 	fmt.Println(arr)
// 	// age := 20
// 	// fmt.Println("Age:", age)
// 	// fmt.Println("Hello, World!")
// 	// printName()
// 	// calculateSum()
// }

// creating files in go

// func createFile(){
// 	fmt.Println("writing file...")
//     file,err := os.Create("test.txt")
// 	if err != nil{
// 		panic(err)
// 	}
// 	length, err := file.WriteString("welcome to golang" +
//         "demonstrates reading and writing operations to a file in golang.")
// 		if err != nil{
// 			panic(err)
// 		}
// 		fmt.Printf("File name: %s", file.Name())
//     fmt.Printf("\nfile length: %d\n", length)
// }

// func readFile(){
// 	fmt.Println("Reading...")
// 	fileName := "test.txt"
// 	data, err := os.ReadFile(fileName)
// 	if err != nil {
//         panic(err)
//     }
// 	fmt.Println("file name " + fileName)
//     fmt.Printf("file size %d\n", len(data))
//     fmt.Printf("file content : %s\n", data)

// }

// build cli

func buildCli(){
	reader := bufio.NewReader(os.Stdin)

    // Ask for filename
    fmt.Print("Enter filename: ")
    filename, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading filename:", err)
        return
    }

    // Remove the newline character from filename
    filename = filename[:len(filename)-1]

    // Ask for file content
    fmt.Print("Enter content for the file: ")
    content, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading content:", err)
        return
    }

    // Create the file
    file, err := os.Create(filename)
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    // Write content to the file
    _, err = file.WriteString(content)
    if err != nil {
        fmt.Println("Error writing to file:", err)
        return
    }

    fmt.Println("File created successfully:", filename)
}

func main() {
    // createFile()
	// readFile()
    buildCli()
}