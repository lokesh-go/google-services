package google

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	configModule "github.com/lokesh-go/google-services/src/config"
	gdriveScopes "github.com/lokesh-go/google-services/src/google/services/drive/scopes"
	utils "github.com/lokesh-go/google-services/src/utils"
)

// Module ...
type Module struct {
	Config *configModule.Config
}

// New ...
func New(config *configModule.Config) *Module {
	// Returns
	return &Module{
		Config: config,
	}
}

// GetClient ...
func (m *Module) GetClient() (client *http.Client, err error) {
	// Gets oauth config
	oauthConfig, err := m.getOAuthConfig()
	if err != nil {
		return nil, err
	}

	// Gets client
	client, err = m.getClient(oauthConfig)
	if err != nil {
		return nil, err
	}

	// Returns
	return client, nil
}

func (m *Module) getOAuthConfig() (oauthConfig *oauth2.Config, err error) {
	// Reads credential file
	bytes, err := ioutil.ReadFile(m.Config.GDrive.Credential.ClientSecret)
	if err != nil {
		return nil, err
	}

	// Gets scops
	scopes := getScopes()

	// Gets google oauth config
	oauthConfig, err = google.ConfigFromJSON(bytes, scopes)
	if err != nil {
		return nil, err
	}

	// Returns
	return oauthConfig, nil
}

func (m *Module) getClient(oauthConfig *oauth2.Config) (client *http.Client, err error) {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	oauthToken, err := m.tokenFromFile()
	if err != nil {
		// Request a token from the web, then returns the retrieved token.
		oauthToken, err = getTokenFromWeb(oauthConfig)
		if err != nil {
			return nil, err
		}

		// Saves token
		err = m.saveToken(oauthToken)
		if err != nil {
			return nil, err
		}
	}

	// Returns
	return oauthConfig.Client(context.Background(), oauthToken), nil
}

func (m *Module) tokenFromFile() (oauthToken *oauth2.Token, err error) {
	// Reads token file
	err = utils.ReadJSONFile(m.Config.GDrive.Credential.Token, &oauthToken)
	if err != nil {
		return nil, err
	}

	// Returns
	return oauthToken, nil
}

func getTokenFromWeb(oauthConfig *oauth2.Config) (oauthToken *oauth2.Token, err error) {
	// Tokens from web
	authURL := oauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		return nil, err
	}

	oauthToken, err = oauthConfig.Exchange(context.TODO(), authCode)
	if err != nil {
		return nil, err
	}

	// Returns
	return oauthToken, nil
}

func (m *Module) saveToken(oauthToken *oauth2.Token) (err error) {
	// Saves token to the path
	file, err := os.OpenFile(m.Config.GDrive.Credential.Token, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	json.NewEncoder(file).Encode(oauthToken)

	// Returns
	return nil
}

func getScopes() (scope string) {
	// Gets drive scopes
	gdriveScope := gdriveScopes.Get()

	// Returns
	return gdriveScope
}
