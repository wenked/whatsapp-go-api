package main

import (
	"fmt"
	"log"
	"whatsapp-go-api/internal/server"
	internal "whatsapp-go-api/internal/wbots"
)

func main() {

	server := server.NewServer()
	log.Println("Starting server on port 8080")
	internal.StartUp()
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server:", err)
		panic("cannot start server")
	}
}
