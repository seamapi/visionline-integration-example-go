package main

import (
	"context"
	"log"

	api "github.com/seamapi/go"
	goclient "github.com/seamapi/go/client"
)

func main() {

	log.Println("Visionline integration example init")

	client := goclient.NewClient(goclient.WithApiKey("seam_test2US6_9G4L2sJPeso5pitYJFa2Jpto"))

	webview, wvErr := client.ConnectWebviews.Create(context.Background(), &api.ConnectWebviewsCreateRequest{
		AcceptedProviders: []api.AcceptedProvider{"visionline"},
	})

	if wvErr != nil {
		log.Panic(wvErr)
	}

	log.Println(webview)

	// open webview.Url and connect the account through webview

}
