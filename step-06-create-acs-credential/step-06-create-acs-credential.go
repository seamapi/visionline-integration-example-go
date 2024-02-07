package main

import (
	"context"
	"log"
	"os"
	"time"

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

	var visionlineSystem *api.AcsSystem

	for _, s := range systems.AcsSystems {
		if s.ExternalType == "visionline_system" {
			visionlineSystem = s
		}
	}

	users, err := client.Acs.Users.List(context.Background(), &acs.UsersListRequest{
		AcsSystemId: &visionlineSystem.AcsSystemId,
	})

	isMultiPhoneSyncCredential := true

	log.Println(users, err)

	startsAt := time.Now().Format("2006-01-02T15:04:00Z")

	endsAt := time.Now().Add(24 * time.Hour).Format("2006-01-02T15:04:00Z")

	log.Println(startsAt)

	credentials, uErr := client.Acs.Credentials.Create(context.Background(), &acs.CredentialsCreateRequest{
		AcsUserId:                  users.AcsUsers[0].AcsUserId,
		AccessMethod:               "mobile_key",
		IsMultiPhoneSyncCredential: &isMultiPhoneSyncCredential,
		StartsAt:                   &startsAt,
		EndsAt:                     &endsAt,
		VisionlineMetadata: &acs.CredentialsCreateRequestVisionlineMetadata{
			CardFormat: acs.CredentialsCreateRequestVisionlineMetadataCardFormatRfid48.Ptr(),
		},
	})

	if uErr != nil {
		log.Panic(uErr)
	}

	log.Println(credentials)
}
