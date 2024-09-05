package main

import (
    "net/http"
	"log"
	"fmt"
	"encoding/json"
)

// * sign is a pointer

func helloHandler(res http.ResponseWriter, req *http.Request){
  if req.URL.Path != "/hello"{
	http.Error(res,"404 not found",http.StatusNotFound)
	return
  }
  if req.Method != "GET"{
	http.Error(res,"GET not supported",http.StatusNotFound)
	return
  }
  
  response := map[string]string{
	"message": "Hello, World!",
	"status": "200",
}

res.Header().Set("Content-Type", "application/json")

if err := json.NewEncoder(res).Encode(response); err != nil {
	http.Error(res, "Error encoding JSON", http.StatusInternalServerError)
	return
}
  fmt.Println(res,"Hello")
}

func formHandler(res http.ResponseWriter, req *http.Request) {
	// Parse form data
	err := req.ParseForm()
	if err != nil {
		http.Error(res, "Error parsing form data", http.StatusInternalServerError)
		return
	}

	// Retrieve form values
	name := req.FormValue("name")
	address := req.FormValue("address")
	email := req.FormValue("email")

	// Write response
	fmt.Fprintf(res, "Name: %s\nAddress: %s\nEmail: %s\n", name, address, email)
}


func main(){
	fileServer :=  http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)
	fmt.Println("STarting the server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}