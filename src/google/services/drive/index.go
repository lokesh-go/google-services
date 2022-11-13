package drive

import (
	"context"
	"errors"
	"net/http"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"

	configModule "github.com/lokesh-go/google-services/src/config"
)

// Service ...
type Serivce struct {
	drive  *drive.Service
	config *configModule.Config
}

// NewService ...
func NewService(client *http.Client, config *configModule.Config) (*Serivce, error) {
	// Gets new drive service
	service, err := drive.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}

	// Returns
	return &Serivce{
		drive:  service,
		config: config,
	}, nil
}

func (s *Serivce) List() (driveDetails []DriveList, err error) {
	// Gets drive list
	driveList, err := s.drive.Drives.List().PageSize(s.config.GDrive.List.PageSize).Do()
	if err != nil {
		return nil, err
	}

	// Checks
	if len(driveList.Drives) == 0 {
		return nil, errors.New("drive list found empty")
	}

	// Range
	driveDetails = []DriveList{}
	for _, d := range driveList.Drives {
		
		// Prepare list
		list := DriveList{
			Id:   d.Id,
			Name: d.Name,
		}

		// Appends
		driveDetails = append(driveDetails, list)
	}

	// Returns
	return driveDetails, nil
}
