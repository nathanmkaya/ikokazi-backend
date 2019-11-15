package model

type Config struct {
	TwitterConsumerKey    string `env:"CONSUMER_KEY"`
	TwitterConsumerSecret string `env:"CONSUMER_SECRET"`
	TwitterAccessToken    string `env:"ACCESS_TOKEN"`
	TwitterAccessSecret   string `env:"ACCESS_SECRET"`
	ProjectID             string `env:"PROJECT_ID"`
	AlgoliaAppID          string `env:"ALGOLIA_APP_ID"`
	AlgoliaAPIKey         string `env:"ALGOLIA_API_KEY"`
}
