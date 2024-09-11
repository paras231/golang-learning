package main

import "fmt"
import (
	"os"
	"bufio"
    // "time"
	"strings"
	"path/filepath"
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

// func buildCli(){
// 	reader := bufio.NewReader(os.Stdin)

//     // Ask for filename
//     fmt.Print("Enter filename: ")
//     filename, err := reader.ReadString('\n')
//     if err != nil {
//         fmt.Println("Error reading filename:", err)
//         return
//     }

//     // Remove the newline character from filename
//     filename = filename[:len(filename)-1]

//     // Ask for file content
//     fmt.Print("Enter content for the file: ")
//     content, err := reader.ReadString('\n')
//     if err != nil {
//         fmt.Println("Error reading content:", err)
//         return
//     }

//     // Create the file
//     file, err := os.Create(filename)
//     if err != nil {
//         fmt.Println("Error creating file:", err)
//         return
//     }
//     defer file.Close()

//     // Write content to the file
//     _, err = file.WriteString(content)
//     if err != nil {
//         fmt.Println("Error writing to file:", err)
//         return
//     }

//     fmt.Println("File created successfully:", filename)
// }

func checkDirExist(path string)bool{
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
        return false
    }
    return info.IsDir()
}

// calculate directory size

func calculateDirectorySize(dirPath string) (int64, error) {
	var totalSize int64 = 0

   err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error)error{
	 if err != nil {
		return err
	 }
	 if !info.IsDir() {
        totalSize += info.Size()
	 }
	 return nil	
   })
   if err != nil {
	return 0, err
}
return totalSize, nil
}

// directory size analyzer

func directorySizeAnalyzer(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter directory name: ")

	// Read the directory name from input
	dirName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Trim the newline character from the input
	dirName = strings.TrimSpace(dirName)

	// Check if the directory exists
	if checkDirExist(dirName) {
		size,err := calculateDirectorySize(dirName)
		if err != nil{
			fmt.Println(err)
		}
		fmt.Printf("Total size of directory '%s': %.2f MB\n", dirName, float64(size)/(1024*1024))
		fmt.Println("Directory exists")
	} else {
		fmt.Println("Directory does not exist")
	}
   
}


// slices in go

// it takes three args . first is data type , second is length , third is capacity


// func sliceUsage(){
//      // mySlice :=  make([]int , 5 , 10)

//     // fmt.Println(mySlice)
//     // mySlice1 := append(mySlice,100,200,300)
//     // it will append initial 5 zeros
//     // fmt.Println(mySlice1)

//     // to avoid initial zeros
//     sliceNew :=  make([]int,0,20)
    
//     updatedSlice  := append(sliceNew,20,30,40)
//     fmt.Println(updatedSlice)

//     // iterating through slcie using for loop

//     for i:=0; i < len(updatedSlice); i++ {
//         fmt.Println(updatedSlice[i])
//     }
//     start := time.Now()
//     for i:=0; i<1000000; i++{
//         fmt.Println(i)
//     }
//     elapsed := time.Since(start).Milliseconds()
//     fmt.Printf("Loop completed in %d ms\n", elapsed)
// }


// goroutines

// let's create goroutines

// func createGoroutines(){
//     fmt.Println("Creating goroutines")
// }

// func numbers(){
//     for i := 0; i<20;i++{
//         time.Sleep(100 * time.Millisecond)
//         fmt.Printf("%d ", i)
//     }
// }

// channels in go
// channels are used to communicate between goroutines

// creating channel


// func worker(ch chan string){
//   time.Sleep(2* time.Second)
//   ch  <- "Data processed"  // send data to channel

// }


// Buffered channels


// go channels with select

// func sendData(ch1, ch2 chan string){
// 	time.Sleep(3*time.Second)
// 	ch1 <- "Data for ch 1" 
// 	time.Sleep(3*time.Second)
// 	ch2 <- "Data for ch 2"

// }

//! worker pool with channels

// func worker(id int ,tasks <-chan int, results chan<- int){
// 	for task := range tasks {
// 		fmt.Printf("Worker %d processing task %d\n", id, task)
// 		// time.Sleep(time.Second) // Simulate work
// 		results <- task * 2      // Send result back
// 	}
// }

func main() {
    // createFile()
	// readFile()
    // buildCli()
    // go createGoroutines()
    // time.Sleep(1 * time.Second)
    // go numbers()
    // time.Sleep(1000 * time.Millisecond)
	// directorySizeAnalyzer()
    // fmt.Println("main function")
	// messages := make(chan string)
	// go func() { messages <- "ping" }()
	// msg := <-messages
    // fmt.Println(msg)
	// ch := make(chan string)
	// go worker(ch) // start worker goroutine
	// fmt.Println("Waiting for worker...")
	// message := <-ch // Receive data from channel
	// fmt.Println(message)

	// ch := make(chan int , 3)
	
	// ch<-1
	// ch <- 2
	// ch <- 3
	// fmt.Println(<-ch) // Outputs 1
	// fmt.Println(<-ch) // Outputs 2
	// fmt.Println(<-ch) // Outputs 3
    
	// ch1 := make(chan string)
	// ch2 := make(chan string)
	// go sendData(ch1, ch2)
	
	// for i :=0; i<2; i++ {
	// 	select{
	// 	case msg1 := <-ch1:
	// 		fmt.Println("Received from channel 1",msg1)
	// 	case msg2 := <-ch2:
	// 		fmt.Println("Received from ch2:", msg2)	
	// 	}
	// }
	// numOfWorkers := 3
	// tasks := make(chan int , 200)
	// results := make(chan int, 200) // Results channel

	// // create a pool of workers
	// for i := 0; i < numOfWorkers; i++ {
	// 	go worker(i,tasks,results)
	// }

	// for t := 1; t <= 200; t++ {
	// 	tasks <- t
	// }
	// close(tasks) // No more tasks will be sent

	// // Collect results
	// for r := 1; r <= 200; r++ {
	// 	fmt.Println("Result:", <-results)
	// }

}