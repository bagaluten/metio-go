/*
 * Copyright 2024 Bagaluten GmbH <contact@bagaluten.email>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package client

import (
	"encoding/json"
	"fmt"

	"github.com/bagaluten/metio-go/types"
	"github.com/nats-io/nats.go"
)

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

// Close closes the connection to the server.
func (c *Client) Close() {
	c.client.Close()
}

// Publish sends the given data to the server.
func (c *Client) Publish(subject string, data []types.Event) error {
	if c.prefix != "" {
		subject = fmt.Sprintf("%s.%s", c.prefix, subject)
	}
	PartialError := PartialError{}
	for index, event := range data {
		bytes, err := json.Marshal(event)
		if err != nil {
			PartialError.addError(err, index)
			continue
		}
		err = c.client.Publish(subject, bytes)
		if err != nil {
			PartialError.addError(err, index)
			continue
		}
	}
	c.client.Flush()

	if PartialError.isEmpty() {
		return nil
	}
	return PartialError

}
