package main

import (
	"context"
	"log"
	"os"

	goclient "github.com/seamapi/go/client"
	"github.com/seamapi/go/useridentities"
)

func main() {

	client := goclient.NewClient(goclient.WithApiKey(os.Getenv("SEAM_API_KEY")))

	janeEmail := "jane@example.com"

	user, uErr := client.UserIdentities.Create(context.Background(), &useridentities.UserIdentitiesCreateRequest{
		EmailAddress: &janeEmail,
	})

	if uErr != nil {
		log.Panic(uErr)
	}

	log.Println(user)
}
