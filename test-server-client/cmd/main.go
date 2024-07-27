package main

import (
	"log"

	"github.com/avGenie/gophkeeper/test-server-client/internal/entity"
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

	// err = client.SaveLoginPasswordUser("google creds", "looogggin", "qwerty", "", token)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = client.SaveText("google text", "tezxt adsad text", "", token)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = client.SaveCard(
	// 	"card 1",
	// 	"111122223333",
	// 	"",
	// 	5,
	// 	2030,
	// 	444,
	// 	"",
	// 	token,
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// loginPass, err := client.GetLoginPasswordUser("google creds", token)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(loginPass)

	// text, err := client.GetTextData("google text", token)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(text)

	// card, err := client.GetCardData("card 1", token)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(card)

	err = client.DeleteDataObject(entity.DataRequestLoginPassword, "google creds", token)
	if err != nil {
		log.Fatal(err)
	}

	// err = client.DeleteDataObject(entity.DataRequestText, "google text", token)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = client.DeleteDataObject(entity.DataRequestCard, "card 1", token)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	client.Stop()
}
