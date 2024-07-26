package main

import (
	"log"

	"github.com/avGenie/gophkeeper/test-server-client/internal/grpc"
)

func main() {
	client := grpc.CreateClient()

	// err := client.RegisterUser("login", "password")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	token, err := client.AuthenticateUser("login", "password")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(token)

	client.Stop()
}
