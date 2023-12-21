package main

import (
	"context"
	"log"

	goclient "github.com/seamapi/go/client"
	"github.com/seamapi/go/useridentities"
)

func main() {

	client := goclient.NewClient(goclient.WithApiKey("seam_test2US6_9G4L2sJPeso5pitYJFa2Jpto"))

	janeEmail := "seam@example.com"

	user, uErr := client.UserIdentities.Create(context.Background(), &useridentities.UserIdentitiesCreateRequest{
		EmailAddress: &janeEmail,
	})

	if uErr != nil {
		log.Panic(uErr)
	}

	log.Println(user)
}
