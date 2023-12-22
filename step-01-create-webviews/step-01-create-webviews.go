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

	seamBridgeWebview, err := client.ConnectWebviews.Create(context.Background(), &api.ConnectWebviewsCreateRequest{
		AcceptedProviders: []api.AcceptedProvider{"seam_bridge"},
	})

	if err != nil {
		log.Panic(err)
	}

	assaAbloyCredentialsWebview, err := client.ConnectWebviews.Create(context.Background(), &api.ConnectWebviewsCreateRequest{
		AcceptedProviders: []api.AcceptedProvider{"assa_abloy_credential_service"},
	})

	if err != nil {
		log.Panic(err)
	}

	visionlineWebview, err := client.ConnectWebviews.Create(context.Background(), &api.ConnectWebviewsCreateRequest{
		AcceptedProviders: []api.AcceptedProvider{"visionline"},
	})

	if err != nil {
		log.Panic(err)
	}

	log.Println("Seam Bridge Webview:")
	log.Println(seamBridgeWebview.Url)
	log.Println("Assa Abloy Credentials Webview:")
	log.Println(assaAbloyCredentialsWebview.Url)
	log.Println("Visionline Webview:")
	log.Println(visionlineWebview.Url)
}
