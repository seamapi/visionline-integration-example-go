package main

import (
	"context"
	"log"
	"os"

	seamapigo "github.com/seamapi/go"
	goclient "github.com/seamapi/go/client"
)

func main() {

	client := goclient.NewClient(goclient.WithApiKey(os.Getenv("SEAM_API_KEY")))

	janeEmail := "jane@example.com"

	user, uErr := client.UserIdentities.Create(context.Background(), &seamapigo.UserIdentitiesCreateRequest{
		EmailAddress: &janeEmail,
	})

	if uErr != nil {
		log.Panic(uErr)
	}

	log.Println(user)
}
