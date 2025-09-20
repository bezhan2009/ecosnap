package configs

import (
	"ecosnap/internal/app/models"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/oauth2"
	"google.golang.org/api/drive/v3"
	"log"
	"os"
)

func ReadSettings() (models.Configs, error) {
	var AppSettings models.Configs

	configFile, err := os.Open("configs/configs.json")
	if err != nil {
		configFile, err = os.Open("configs/example.json")
		if err != nil {
			return models.Configs{}, errors.New(fmt.Sprintf("Couldn't open config file. Error is: %s", err.Error()))
		}
	}

	defer func(configFile *os.File) {
		err = configFile.Close()
		if err != nil {
			log.Fatal("Couldn't close config file. Error is: ", err.Error())
		}
	}(configFile)

	if err = json.NewDecoder(configFile).Decode(&AppSettings); err != nil {
		return models.Configs{}, errors.New(fmt.Sprintf("Couldn't decode settings json file. Error is: %s", err.Error()))
	}

	AppSettings.OAuth2 = oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT"),
		Scopes:       []string{drive.DriveScope},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://oauth2.googleapis.com/token",
		},
	}

	return AppSettings, nil
}
