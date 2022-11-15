package config

// Config ...
type Config struct {
	GDrive struct {
		Credential struct {
			ClientSecret string
			Token        string
		}
		FileSearch struct {
			PageSize                  int64
			Extend                    bool
			SupportsAllDrives         bool
			IncludeItemsFromAllDrives bool
			Corpora                   string
			Spaces                    string
			Fields                    string
			Query                     struct {
				FileContains      string
				NotContainsFolder string
				NotContainsTrash  string
			}
		}
	}
}
