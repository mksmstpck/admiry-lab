package conectors

import (
	"github.com/labstack/gommon/log"
	"github.com/nats-io/nats.go"
)

func NewNats(url string) (*nats.EncodedConn, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	log.Info("conect to nats server")
	return c, nil
}
