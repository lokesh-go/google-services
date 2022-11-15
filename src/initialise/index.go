package initialise

import (
	"fmt"

	configModule "github.com/lokesh-go/google-services/src/config"
	google "github.com/lokesh-go/google-services/src/google"
	drive "github.com/lokesh-go/google-services/src/google/services/drive"
	utils "github.com/lokesh-go/google-services/src/utils"
)

// Init ...
func Init() (err error) {
	// Initialize config
	config, err := initConfig()
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
	driveService, err := drive.NewService(client, config)
	if err != nil {
		return err
	}

	// Gets list
	list, err := driveService.FileSearch("endgame imax 7.1")
	if err != nil {
		return err
	}

	// Ranging list
	for k, v := range list {
		fmt.Println(k+1, " --- ", v.Name, " --- ", v.MimeType)
	}

	// Returns
	return nil
}

func initConfig() (config *configModule.Config, err error) {
	// Binds config models
	err = utils.ReadJSONFile("config/config.json", &config)
	if err != nil {
		return nil, err
	}

	// Returns
	return config, nil
}
