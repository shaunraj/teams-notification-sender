package main

import (
	"context"
	"fmt"
	"teams/reaction-poster/azure"
	"teams/reaction-poster/config"

	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

func main() {
	config := config.GetApplicationConfig()
	azureAuthCredential := azure.AuthenticateToAzure(config)
	graphClient := azure.GetGraphClient(azureAuthCredential)

	result, err := graphClient.Me().Get(context.Background(), nil)

	if err != nil {
		printOdataError(err)
	}
	println(result.GetAboutMe())
}

func printOdataError(err error) {
	switch err.(type) {
	case *odataerrors.ODataError:
		typed := err.(*odataerrors.ODataError)
		fmt.Printf("error: %s\n", typed.Error())
		if terr := typed.GetError(); terr != nil {
			fmt.Printf("code: %s\n", *terr.GetCode())
			fmt.Printf("msg: %s\n", *terr.GetMessage())
		}
	default:
		fmt.Printf("%T > error: %#v\n", err, err)
	}
}
