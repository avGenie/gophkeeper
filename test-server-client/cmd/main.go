package main

import (
	"fmt"
	"log"

	"github.com/avGenie/gophkeeper/test-server-client/internal/grpc"
)

func main() {
	client := grpc.CreateClient()

	// err := client.RegisterUser("login", "password")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// token, err := client.AuthenticateUser("login", "password")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(token)

	// err := client.SaveLoginPasswordUser("google creds", "looogggin", "qwerty", "", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjIwNzk1OTksIlVzZXJJRCI6ImM5MjhjYTM4LWVkZDMtNGUxOS1iNDBiLTZhOTljMGExMWM0MCJ9.0t3U5_Rt5jImqO032M1vbMRWnxwTcy_nFFW0hzCcLiw")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err := client.SaveText("google text", "tezxt adsad text", "", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjIwNzk1OTksIlVzZXJJRCI6ImM5MjhjYTM4LWVkZDMtNGUxOS1iNDBiLTZhOTljMGExMWM0MCJ9.0t3U5_Rt5jImqO032M1vbMRWnxwTcy_nFFW0hzCcLiw")
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
	// 	"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjIwNzk1OTksIlVzZXJJRCI6ImM5MjhjYTM4LWVkZDMtNGUxOS1iNDBiLTZhOTljMGExMWM0MCJ9.0t3U5_Rt5jImqO032M1vbMRWnxwTcy_nFFW0hzCcLiw",
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

	loginPass, err := client.GetLoginPasswordUser("google creds", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjIwNzk1OTksIlVzZXJJRCI6ImM5MjhjYTM4LWVkZDMtNGUxOS1iNDBiLTZhOTljMGExMWM0MCJ9.0t3U5_Rt5jImqO032M1vbMRWnxwTcy_nFFW0hzCcLiw")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(loginPass)

	text, err := client.GetTextData("google text", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjIwNzk1OTksIlVzZXJJRCI6ImM5MjhjYTM4LWVkZDMtNGUxOS1iNDBiLTZhOTljMGExMWM0MCJ9.0t3U5_Rt5jImqO032M1vbMRWnxwTcy_nFFW0hzCcLiw")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(text)

	card, err := client.GetCardData("card 1", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjIwNzk1OTksIlVzZXJJRCI6ImM5MjhjYTM4LWVkZDMtNGUxOS1iNDBiLTZhOTljMGExMWM0MCJ9.0t3U5_Rt5jImqO032M1vbMRWnxwTcy_nFFW0hzCcLiw")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(card)

	client.Stop()
}
