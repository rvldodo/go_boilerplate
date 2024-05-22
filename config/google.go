package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleConfig struct {
	GoogleLoginConfig oauth2.Config
}

var (
	AppConfig  GoogleConfig
	GoogleInit = GoogleConfigInit()
)

func GoogleConfigInit() oauth2.Config {
	AppConfig.GoogleLoginConfig = oauth2.Config{
		RedirectURL:  "http://localhost:1337/api/v1/google_callback",
		ClientID:     Envs.GoogleClientID,
		ClientSecret: Envs.GoogleClientSecret,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	return AppConfig.GoogleLoginConfig
}
