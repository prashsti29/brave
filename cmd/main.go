package main

import (
    "fmt"
    "github.com/prashsti29/brave/config"
)

func main() {
    db := config.ConnectDB()
    defer db.Close()
    fmt.Println("Server starting...")
}