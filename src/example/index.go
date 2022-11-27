package example

import (
	"fmt"

	google "github.com/lokesh-go/google-services/src/google"
	drive "github.com/lokesh-go/google-services/src/google/services/drive"
)

// Init ...
func Init() (err error) {
	// Initialize config
	config := getConfig()
	if err != nil {
		return err
	}

	// Gets google client
	googleModule := google.New(config)
	client, err := googleModule.GetClient()
	if err != nil {
		return err
	}

	// Gets drive service
	driveService, err := drive.NewService(client)
	if err != nil {
		return err
	}

	// Gets list
	list, err := driveService.FileSearch("avengers", false, true)
	if err != nil {
		return err
	}

	// Prints list
	for k, file := range list {
		fmt.Println(k+1, " -- ", file.DownloadLink)
	}

	// Returns
	return nil
}

func getConfig() (config *google.Config) {
	return &google.Config{
		ClientSecretFilePath: "/Users/lokeshchandra/Desktop/google/credentials.json",
		TokenPath:            "/Users/lokeshchandra/Desktop/google/token.json",
	}
}
