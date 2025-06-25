package main

import (
	"log"

	"github.com/argon-dev22/sample.gmail-api/app/apis"
	"github.com/argon-dev22/sample.gmail-api/app/auth"
	"github.com/argon-dev22/sample.gmail-api/app/config"
)

func main() {
	env := config.GetEnv()
	srv := auth.GetService(env)
	user := "me"

	// Gmail API 呼び出し
	res := apis.GetProfile(srv, user)

	log.Printf("%+v\n", res)
}
