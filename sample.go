package main

import (
	"context"
	"fmt"

	"github.com/achan-cogito/example-test/user"
)

func main() {
	fmt.Printf("Starting the test\n")
	configuration := user.NewConfiguration()
	apiClient := user.NewAPIClient(configuration)
	// resp, httpRes, err := apiClient.PetAPI.AddPet(context.Background()).Execute()
	resp, httpRes, err := apiClient.UserAPI.GetUserByName(context.Background(), "achan").Execute()
	if err != null {
		fmt.Printf("pass")
	} else {
		fmt.Printf("failed")
	}
}
