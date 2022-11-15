package drive

import (
	"context"
	"net/http"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
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

// FileSearch ...
func (s *Serivce) FileSearch(searchKey string) (fileDetails []FileList, err error) {
	// Forms file search query
	searchQuery := s.formFileSearchQuery(searchKey)

	// Gets file list
	files, err := s.getFileList(searchQuery)
	if err != nil {
		return nil, err
	}

	// Range
	fileDetails = []FileList{}
	for _, f := range files {
		list := FileList{
			Id:            f.Id,
			DriveId:       f.DriveId,
			Name:          f.Name,
			Size:          f.Size,
			FileExtension: f.FileExtension,
			MimeType:      f.MimeType,
			Md5Checksum:   f.Md5Checksum,
		}

		// Appends
		fileDetails = append(fileDetails, list)
	}

	// Returns
	return fileDetails, nil
}

func (s *Serivce) formFileSearchQuery(searchKey string) (query string) {
	// Gets file search config
	fileConfig := s.config.GDrive.FileSearch

	// Forms search query
	query = fileConfig.Query.NotContainsFolder + " and " + fileConfig.Query.FileContains + searchKey + "' and " + fileConfig.Query.NotContainsTrash

	// Returns
	return query
}

func (s *Serivce) getFileList(searchQuery string) (files []*drive.File, err error) {
	var pageToken string
	files = []*drive.File{}
	fileConfig := s.config.GDrive.FileSearch

	for {
		// Gets file lists
		fileList, err := s.drive.Files.List().SupportsAllDrives(fileConfig.SupportsAllDrives).IncludeItemsFromAllDrives(fileConfig.IncludeItemsFromAllDrives).Corpora(fileConfig.Corpora).Spaces(fileConfig.Spaces).Fields(googleapi.Field(fileConfig.Fields)).Q(searchQuery).PageSize(fileConfig.PageSize).PageToken(pageToken).Do()
		if err != nil {
			return nil, err
		}

		// Appends
		files = append(files, fileList.Files...)

		// Searches not extends
		if !fileConfig.Extend {
			break
		}

		// Checks
		pageToken = fileList.NextPageToken
		if pageToken == "" {
			break
		}
	}

	// Returns
	return files, nil
}
