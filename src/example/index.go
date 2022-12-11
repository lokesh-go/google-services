package example

import (
	"fmt"
	"log"

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

	details, err := driveService.DriveDetails()
	if err != nil {
		log.Fatal("Error: ", err.Error())
	}
	fmt.Println(len(details.ShareDriveNames))
	// upload file
	// id, err := driveService.FileCreate("testing.html", "text/html", "xyzabc", "example/a.html")
	// if err != nil {
	// 	log.Fatal("Error: ", err.Error())
	// }

	// fmt.Println("Uploaded File Id: ", id)

	// Gets list
	// list, err := driveService.FileSearch("avengers", false, true)
	// if err != nil {
	// 	return err
	// }

	// // Prints list
	// for k, file := range list {
	// 	fmt.Println(k+1, " -- ", file.DownloadLink)
	// }

	// Returns
	return nil
}

func getConfig() (config *google.Config) {
	return &google.Config{
		ClientSecretFilePath: "/Users/lokeshchandra/Desktop/google/credentials.json",
		TokenPath:            "/Users/lokeshchandra/Desktop/google/token.json",
		Scopes: google.Scopes{
			DriveScope: true,
		},
	}
}
