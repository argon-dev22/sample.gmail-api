package apis

import (
	"log"

	"google.golang.org/api/gmail/v1"
)

func ListLabels(srv *gmail.Service, user string) *gmail.ListLabelsResponse {
	labelsRes, err := srv.Users.Labels.List(user).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve labels: %v", err)
	}
	return labelsRes
}
