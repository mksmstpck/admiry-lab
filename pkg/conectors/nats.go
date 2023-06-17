package conectors

import "github.com/nats-io/nats.go"

func NewNats(url string) (*nats.EncodedConn, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	return c, nil
}
