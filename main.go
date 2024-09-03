package main

import (
	"api-go-gin/database"
	"api-go-gin/routes"
	"fmt"
)

func main() {
	port := ":8000"
	fmt.Println("Server running on port " + port)
	database.ConnectDB()
	routes.HandleRequests(port)
}
