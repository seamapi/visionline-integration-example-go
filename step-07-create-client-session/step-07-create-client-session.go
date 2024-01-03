package main

import (
	"context"
	"log"
	"os"

	api "github.com/seamapi/go"
	goclient "github.com/seamapi/go/client"
)

func main() {
	client := goclient.NewClient(goclient.WithApiKey(os.Getenv("SEAM_API_KEY")))

	userIdentities, uErr := client.UserIdentities.List(context.Background())

	if uErr != nil {
		log.Panic(uErr)
	}

	if len(userIdentities) == 0 {
		log.Panic("No user indentity found")

	}

	clientSession, csErr := client.ClientSessions.Create(context.Background(), &api.ClientSessionsCreateRequest{
		UserIdentifierKey: userIdentities[0].UserIdentityKey,
	})

	if csErr != nil {
		log.Panic(csErr)
	}

	log.Println(clientSession)

}
