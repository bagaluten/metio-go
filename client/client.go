package client

import "github.com/nats-io/nats.go"

type Config struct {
	// The address of the server to connect to.
	Host string
	// The prefix that is applied to names that are used when publishing messages
	Prefix *string
}

// Client is a low-level way to interact with a Metio compliant endpoint.
type Client struct {
	client *nats.Conn
	// The prefix that is applied to names that are used when publishing messages
	prefix string
}

// NewClient creates a new client that connects to the server at the given address.
func NewClient(config Config) (*Client, error) {
	client, err := nats.Connect(config.Host)
	if err != nil {
		return nil, err
	}
	prefix := ""
	if config.Prefix != nil {
		prefix = *config.Prefix
	}
	return &Client{client: client, prefix: prefix}, nil
}
