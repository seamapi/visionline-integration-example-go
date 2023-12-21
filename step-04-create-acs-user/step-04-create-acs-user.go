package main

import (
	"context"
	"log"

	api "github.com/seamapi/go"
	"github.com/seamapi/go/acs"
	goclient "github.com/seamapi/go/client"
)

func main() {

	client := goclient.NewClient(goclient.WithApiKey("seam_test2US6_9G4L2sJPeso5pitYJFa2Jpto"))

	systems, sErr := client.Acs.Systems.List(context.Background(), nil)

	if sErr != nil {
		log.Panic(sErr)
	}

	var visionlineSystem *api.AcsSystem

	for _, s := range systems.AcsSystems {
		if s.ExternalType == "visionline_system" {
			visionlineSystem = s
		}
	}

	users, _ := client.UserIdentities.List(context.Background())

	FullName := "New User Full Name"

	user, uErr := client.Acs.Users.Create(context.Background(), &acs.UsersCreateRequest{
		AcsSystemId:    visionlineSystem.AcsSystemId,
		UserIdentityId: &users[0].UserIdentityId,
		FullName:       &FullName,
	})

	if uErr != nil {
		log.Panic(uErr)
	}

	log.Println(user)
}