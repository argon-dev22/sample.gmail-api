package apis

import (
	"log"

	"google.golang.org/api/gmail/v1"
)

func GetProfile(srv *gmail.Service, user string) *gmail.Profile {
	profileRes, err := srv.Users.GetProfile(user).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve profile: %v", err)
	}
	return profileRes
}
