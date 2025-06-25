package apis

import (
	"fmt"
	"log"

	"github.com/argon-dev22/sample.gmail-api/app/config"
	"google.golang.org/api/gmail/v1"
)

func PostWatch(env *config.Env, srv *gmail.Service, user string) *gmail.WatchResponse {
	watchLabelIds := []string{"INBOX"}
	topicName := fmt.Sprintf("projects/%s/topics/%s", env.GoogleCloudProjectName, env.GoogleCloudWatchTopicName)
	watchRes, err := srv.Users.Watch(user, &gmail.WatchRequest{
		TopicName: topicName,
		LabelIds:  watchLabelIds,
	}).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve watch action: %v", err)
	}

	return watchRes
}
