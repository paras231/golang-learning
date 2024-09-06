package main

import (
    // "net/http"
	"log"
	"fmt"
	// "encoding/json"
	"database/sql"
	 _ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "paras"
    password = "paras123"
    dbname   = "goproject"
)

// * sign is a pointer

// func helloHandler(res http.ResponseWriter, req *http.Request){
//   if req.URL.Path != "/hello"{
// 	http.Error(res,"404 not found",http.StatusNotFound)
// 	return
//   }
//   if req.Method != "GET"{
// 	http.Error(res,"GET not supported",http.StatusNotFound)
// 	return
//   }
  
//   response := map[string]string{
// 	"message": "Hello, World!",
// 	"status": "200",
// }

// res.Header().Set("Content-Type", "application/json")

// if err := json.NewEncoder(res).Encode(response); err != nil {
// 	http.Error(res, "Error encoding JSON", http.StatusInternalServerError)
// 	return
// }
//   fmt.Println(res,"Hello")
// }

// func formHandler(res http.ResponseWriter, req *http.Request) {
// 	// Parse form data
// 	err := req.ParseForm()
// 	if err != nil {
// 		http.Error(res, "Error parsing form data", http.StatusInternalServerError)
// 		return
// 	}

// 	// Retrieve form values
// 	name := req.FormValue("name")
// 	address := req.FormValue("address")
// 	email := req.FormValue("email")

// 	// Write response
// 	fmt.Fprintf(res, "Name: %s\nAddress: %s\nEmail: %s\n", name, address, email)
// }


func main(){
	// fileServer :=  http.FileServer(http.Dir("./static"))
	// http.Handle("/", fileServer)
	// http.HandleFunc("/form",formHandler)
	// http.HandleFunc("/hello",helloHandler)
	// fmt.Println("STarting the server...")
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	log.Fatal(err)
	// }
   
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Successfully connected to PostgreSQL!")
}