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
	list, err := driveService.List()
	if err != nil {
		return err
	}

	fmt.Println(list)

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
