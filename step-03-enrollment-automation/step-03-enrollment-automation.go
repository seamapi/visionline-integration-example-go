package main

import (
	"context"
	"log"
	"os"

	api "github.com/seamapi/go"
	goclient "github.com/seamapi/go/client"
	"github.com/seamapi/go/useridentities"
)

func main() {

	client := goclient.NewClient(goclient.WithApiKey(os.Getenv("SEAM_API_KEY")))

	systems, sErr := client.Acs.Systems.List(context.Background(), nil)

	if sErr != nil {
		log.Panic(sErr)
	}

	log.Println(systems)

	var assaAbloySystem *api.AcsSystem

	for _, s := range systems.AcsSystems {
		if s.ExternalType == "assa_abloy_credential_service_user" {
			assaAbloySystem = s
		}
	}

	users, err := client.UserIdentities.List(context.Background())

	shouldCreateCredentialUser := true

	eaResponse, err := client.UserIdentities.EnrollmentAutomations.Launch(context.Background(), &useridentities.EnrollmentAutomationsLaunchRequest{
		UserIdentityId:               users[0].UserIdentityId,
		CredentialManagerAcsSystemId: assaAbloySystem.AcsSystemId,
		CreateCredentialManagerUser:  &shouldCreateCredentialUser,
	})

	if err != nil {
		log.Panic(err)
	}

	log.Println(eaResponse)
}
