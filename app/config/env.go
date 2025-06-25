package config

import (
	"log"

	"github.com/joeshaw/envdecode"
)

type Env struct {
	GoogleCloudCredPath        string `env:"GOOGLE_CLOUD_CREDENTIAL_PATH,required"`
	GoogleCloudAccessTokenPath string `env:"GOOGLE_CLOUD_ACCESS_TOKEN_PATH,required"`
	GoogleCloudProjectName     string `env:"GOOGLE_CLOUD_PROJECT_NAME"`
	GoogleCloudWatchTopicName  string `env:"GOOGLE_CLOUD_WATCH_TOPIC_NAME"`
}

func GetEnv() *Env {
	var env Env
	if err := envdecode.StrictDecode(&env); err != nil {
		log.Fatalf("Failed to decode environment variables: %v", err)
	}

	return &env
}
