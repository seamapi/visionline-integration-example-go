package main

import (
	"context"
	"log"
	"os"

	api "github.com/seamapi/go"
	"github.com/seamapi/go/acs"
	goclient "github.com/seamapi/go/client"
)

func main() {

	client := goclient.NewClient(goclient.WithApiKey(os.Getenv("SEAM_API_KEY")))

	systems, sErr := client.Acs.Systems.List(context.Background(), nil)

	if sErr != nil {
		log.Panic(sErr)
	}

	if len(systems.AcsSystems) == 0 {
		log.Panic("No acs system found")

	}

	var visionlineSystem *api.AcsSystem

	for _, s := range systems.AcsSystems {
		if s.ExternalType == "visionline_system" {
			visionlineSystem = s
		}
	}

	users, uErr := client.Acs.Users.List(context.Background(), &acs.UsersListRequest{
		AcsSystemId: visionlineSystem.AcsSystemId,
	})

	if uErr != nil {
		log.Panic(uErr)
	}

	if len(users.AcsUsers) == 0 {
		log.Panic("No acs user found")

	}

	entrancesResponse, eErr := client.Acs.Entrances.List(context.TODO(), nil)

	if eErr != nil {
		log.Panic(eErr)
	}

	user, uErr := client.Acs.Entrances.GrantAccess(context.Background(), &acs.EntrancesGrantAccessRequest{
		AcsUserId:     users.AcsUsers[0].AcsUserId,
		AcsEntranceId: entrancesResponse.AcsEntrances[0].AcsEntranceId, //choose the entrance
	})

	if uErr != nil {
		log.Panic(uErr)
	}

	log.Println(user)
}
