package services

import "github.com/labstack/gommon/log"

func NatsPublishError(err error) {
	if err != nil {
		log.Error("NatsPublishError: ", err)
	}
}
