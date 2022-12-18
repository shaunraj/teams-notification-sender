package azure

import (
	"log"
	"teams/reaction-poster/config"

	azidentity "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	auth "github.com/microsoft/kiota-authentication-azure-go"
)

func AuthenticateToAzure(applicationConfig config.AzureConfig) *auth.AzureIdentityAuthenticationProvider {
	credential, err := azidentity.NewInteractiveBrowserCredential(&azidentity.InteractiveBrowserCredentialOptions{
		TenantID:    applicationConfig.TenantId,
		ClientID:    applicationConfig.ClientId,
		RedirectURL: "ms-appx-web://microsoft.aad.brokerplugin/48735012-9e88-4e59-954c-256f7072fe2e",
	})

	if err != nil {
		log.Fatal(err)
	}

	identity, err := auth.NewAzureIdentityAuthenticationProviderWithScopes(credential, []string{"User.Read"})

	if err != nil {
		log.Fatal(err)
	}
	return identity
}
