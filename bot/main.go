package main

import (
	"github.com/nlopes/slack"
	log "github.com/Sirupsen/logrus"
	"os"
)

func init() {
}

func main() {
	log.Info("starting bot");
	api := slack.New(getEnvVar("SLACK_API_KEY"))
	api.SetDebug(false)

	rtm := api.NewRTM()
	go rtm.ManageConnection()


	for {
		select {
		case msg := <-rtm.IncomingEvents:
			log.Debug("incoming event received")

			switch event := msg.Data.(type) {

			case *slack.ConnectedEvent:
				log.Info(event.Info)

			case *slack.MessageEvent:
				log.Info(event)

			case *slack.RTMError:
				log.Info(event)

			case *slack.InvalidAuthEvent:
				log.Info(event)
				break;

			default:

			}
		}
	}

}

func processMessage(client *slack.Client) {
}

func getEnvVar(envVariableName string) string {
	value, _ := os.LookupEnv(envVariableName)
	return value
}
