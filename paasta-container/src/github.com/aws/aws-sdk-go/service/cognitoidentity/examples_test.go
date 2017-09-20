// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cognitoidentity_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentity"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleCognitoIdentity_CreateIdentityPool() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.CreateIdentityPoolInput{
		AllowUnauthenticatedIdentities: aws.Bool(true),                 // Required
		IdentityPoolName:               aws.String("IdentityPoolName"), // Required
		DeveloperProviderName:          aws.String("DeveloperProviderName"),
		OpenIdConnectProviderARNs: []*string{
			aws.String("ARNString"), // Required
			// More values...
		},
		SupportedLoginProviders: map[string]*string{
			"Key": aws.String("IdentityProviderId"), // Required
			// More values...
		},
	}
	resp, err := svc.CreateIdentityPool(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentity_DeleteIdentities() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.DeleteIdentitiesInput{
		IdentityIdsToDelete: []*string{ // Required
			aws.String("IdentityId"), // Required
			// More values...
		},
	}
	resp, err := svc.DeleteIdentities(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentity_DeleteIdentityPool() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.DeleteIdentityPoolInput{
		IdentityPoolId: aws.String("IdentityPoolId"), // Required
	}
	resp, err := svc.DeleteIdentityPool(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentity_DescribeIdentity() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.DescribeIdentityInput{
		IdentityId: aws.String("IdentityId"), // Required
	}
	resp, err := svc.DescribeIdentity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentity_DescribeIdentityPool() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.DescribeIdentityPoolInput{
		IdentityPoolId: aws.String("IdentityPoolId"), // Required
	}
	resp, err := svc.DescribeIdentityPool(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentity_GetCredentialsForIdentity() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.GetCredentialsForIdentityInput{
		IdentityId: aws.String("IdentityId"), // Required
		Logins: map[string]*string{
			"Key": aws.String("IdentityProviderToken"), // Required
			// More values...
		},
	}
	resp, err := svc.GetCredentialsForIdentity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentity_GetId() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.GetIdInput{
		IdentityPoolId: aws.String("IdentityPoolId"), // Required
		AccountId:      aws.String("AccountId"),
		Logins: map[string]*string{
			"Key": aws.String("IdentityProviderToken"), // Required
			// More values...
		},
	}
	resp, err := svc.GetId(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentity_GetIdentityPoolRoles() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.GetIdentityPoolRolesInput{
		IdentityPoolId: aws.String("IdentityPoolId"), // Required
	}
	resp, err := svc.GetIdentityPoolRoles(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentity_GetOpenIdToken() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.GetOpenIdTokenInput{
		IdentityId: aws.String("IdentityId"), // Required
		Logins: map[string]*string{
			"Key": aws.String("IdentityProviderToken"), // Required
			// More values...
		},
	}
	resp, err := svc.GetOpenIdToken(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentity_GetOpenIdTokenForDeveloperIdentity() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput{
		IdentityPoolId: aws.String("IdentityPoolId"), // Required
		Logins: map[string]*string{ // Required
			"Key": aws.String("IdentityProviderToken"), // Required
			// More values...
		},
		IdentityId:    aws.String("IdentityId"),
		TokenDuration: aws.Int64(1),
	}
	resp, err := svc.GetOpenIdTokenForDeveloperIdentity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentity_ListIdentities() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.ListIdentitiesInput{
		IdentityPoolId: aws.String("IdentityPoolId"), // Required
		MaxResults:     aws.Int64(1),                 // Required
		HideDisabled:   aws.Bool(true),
		NextToken:      aws.String("PaginationKey"),
	}
	resp, err := svc.ListIdentities(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentity_ListIdentityPools() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.ListIdentityPoolsInput{
		MaxResults: aws.Int64(1), // Required
		NextToken:  aws.String("PaginationKey"),
	}
	resp, err := svc.ListIdentityPools(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentity_LookupDeveloperIdentity() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.LookupDeveloperIdentityInput{
		IdentityPoolId:          aws.String("IdentityPoolId"), // Required
		DeveloperUserIdentifier: aws.String("DeveloperUserIdentifier"),
		IdentityId:              aws.String("IdentityId"),
		MaxResults:              aws.Int64(1),
		NextToken:               aws.String("PaginationKey"),
	}
	resp, err := svc.LookupDeveloperIdentity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentity_MergeDeveloperIdentities() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.MergeDeveloperIdentitiesInput{
		DestinationUserIdentifier: aws.String("DeveloperUserIdentifier"), // Required
		DeveloperProviderName:     aws.String("DeveloperProviderName"),   // Required
		IdentityPoolId:            aws.String("IdentityPoolId"),          // Required
		SourceUserIdentifier:      aws.String("DeveloperUserIdentifier"), // Required
	}
	resp, err := svc.MergeDeveloperIdentities(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentity_SetIdentityPoolRoles() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.SetIdentityPoolRolesInput{
		IdentityPoolId: aws.String("IdentityPoolId"), // Required
		Roles: map[string]*string{ // Required
			"Key": aws.String("ARNString"), // Required
			// More values...
		},
	}
	resp, err := svc.SetIdentityPoolRoles(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentity_UnlinkDeveloperIdentity() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.UnlinkDeveloperIdentityInput{
		DeveloperProviderName:   aws.String("DeveloperProviderName"),   // Required
		DeveloperUserIdentifier: aws.String("DeveloperUserIdentifier"), // Required
		IdentityId:              aws.String("IdentityId"),              // Required
		IdentityPoolId:          aws.String("IdentityPoolId"),          // Required
	}
	resp, err := svc.UnlinkDeveloperIdentity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentity_UnlinkIdentity() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.UnlinkIdentityInput{
		IdentityId: aws.String("IdentityId"), // Required
		Logins: map[string]*string{ // Required
			"Key": aws.String("IdentityProviderToken"), // Required
			// More values...
		},
		LoginsToRemove: []*string{ // Required
			aws.String("IdentityProviderName"), // Required
			// More values...
		},
	}
	resp, err := svc.UnlinkIdentity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCognitoIdentity_UpdateIdentityPool() {
	svc := cognitoidentity.New(nil)

	params := &cognitoidentity.IdentityPool{
		AllowUnauthenticatedIdentities: aws.Bool(true),                 // Required
		IdentityPoolId:                 aws.String("IdentityPoolId"),   // Required
		IdentityPoolName:               aws.String("IdentityPoolName"), // Required
		DeveloperProviderName:          aws.String("DeveloperProviderName"),
		OpenIdConnectProviderARNs: []*string{
			aws.String("ARNString"), // Required
			// More values...
		},
		SupportedLoginProviders: map[string]*string{
			"Key": aws.String("IdentityProviderId"), // Required
			// More values...
		},
	}
	resp, err := svc.UpdateIdentityPool(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
