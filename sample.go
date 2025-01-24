package main

import (
	"context"
	"github.com/achan-cogito/example-test/pkg"
	"fmt"
)

func main() {
	fmt.Printf("Starting the test\n")
	configuration := pkg.NewConfiguration()
	apiClient := pkg.NewAPIClient(configuration)
	// resp, httpRes, err := apiClient.PetAPI.AddPet(context.Background()).Execute()
	resp, httpRes, err := apiClient.UserAPI.GetUserByName(context.Background(), "achan").Execute()
	if err != null {
		fmt.Printf("pass")
	} else {
		fmt.Printf("failed")
	}
}
