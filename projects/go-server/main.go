package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "strings"

    _ "github.com/lib/pq"
)

// Database struct to manage database connection and operations
type Database struct {
    Conn *sql.DB
}

// User struct to handle incoming user data
type User struct {
    Name string `json:"name"`
    Age  int `json:"age"`
    ID int `json:"id"`
}

// Response struct to handle the response
type Response struct {
    ID      int    `json:"id"`
    Message string `json:"message"`
}

// single user struct response

type UserResponse struct {
    ID int `json:"id"`
    Age int `json:"age"`
    Name string `json:name"`
}

// InitializeDatabase initializes the database connection and creates the table
func (db *Database) InitializeDatabase() error {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    var err error
    db.Conn, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        return fmt.Errorf("error opening database connection: %w", err)
    }

    err = db.Conn.Ping()
    if err != nil {
        return fmt.Errorf("error pinging database: %w", err)
    }

    _, err = db.Conn.Exec(`
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name TEXT,
        age INT
    )`)
    if err != nil {
        return fmt.Errorf("error creating table: %w", err)
    }

    return nil
}

// Close closes the database connection
func (db *Database) Close() {
    if db.Conn != nil {
        db.Conn.Close()
    }
}

// CreateUser inserts a new user into the database
func (db *Database) CreateUser(w http.ResponseWriter, r *http.Request) {
    var user User

    // Parse the JSON body request
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Convert the age from string to int
    // ageInt, err := strconv.Atoi(user.Age)
    // if err != nil {
    //     http.Error(w, "Invalid age format", http.StatusBadRequest)
    //     return
    // }

    sqlStatement := `INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id`
    var userId int

    // Execute the statement and capture the returned id
    err = db.Conn.QueryRow(sqlStatement, user.Name, user.Age).Scan(&userId)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Println("User created")

    // Return success response
    response := Response{
        ID:      userId,
        Message: "User successfully created",
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

// get a user by id from db

func (db *Database) getUserById(w http.ResponseWriter, r *http.Request) {
    var user User
    // Extract ID from URL path
    parts := strings.Split(r.URL.Path, "/")
    if len(parts) < 3 {
        http.Error(w, "Invalid request path", http.StatusBadRequest)
        return
    }

    userId := parts[2] // Extract ID from path
    idInt, err := strconv.Atoi(userId)
    if err != nil {
        http.Error(w, "Invalid user ID format", http.StatusBadRequest)
        return
    }

    // SQL statement to select user by ID
    sqlStatement := `SELECT id, name, age FROM users WHERE id = $1`

    // Query the database and scan the result into the user struct fields
    row := db.Conn.QueryRow(sqlStatement, idInt)
    err = row.Scan(&user.ID, &user.Name, &user.Age)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "No user found", http.StatusNotFound)
            return
        }
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Create a response struct and encode it to JSON
    response := UserResponse{
        ID:   user.ID,
        Name: user.Name,
        Age:  user.Age,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}


const (
    host     = "localhost"
    port     = 5432
    user     = "paras"
    password = "paras123"
    dbname   = "goproject"
)

func main() {
    db := &Database{}

    // Initialize database
    if err := db.InitializeDatabase(); err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    http.HandleFunc("/create-user", db.CreateUser)
    http.HandleFunc("/user/",db.getUserById)
    fmt.Println("Starting the server...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("Error starting server: ", err)
    }
}
