package azure

import (
	"log"

	auth "github.com/microsoft/kiota-authentication-azure-go"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
)

func GetGraphClient(credential *auth.AzureIdentityAuthenticationProvider) *msgraphsdkgo.GraphServiceClient {
	adapter, err := msgraphsdkgo.NewGraphRequestAdapter(credential)

	if err != nil {
		log.Fatal(err)
	}

	return msgraphsdkgo.NewGraphServiceClient(adapter)
}
