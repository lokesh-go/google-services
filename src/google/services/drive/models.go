package drive

// FileList ...
type FileList struct {
	Id            string `json:"id"`
	DriveId       string `json:"driveId"`
	Name          string `json:"name"`
	Size          int64  `json:"size"`
	FileExtension string `json:"fileExtension"`
	MimeType      string `json:"mimeType"`
	Md5Checksum   string `json:"md5Checksum"`
}
