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
	"context"
	"encoding/json"
	"fmt"

	"github.com/bagaluten/metio-go/types"
	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel/trace"
)

type Config struct {
	// The address of the server to connect to.
	Host string
	// The prefix that is applied to names that are used when publishing messages
	Prefix *string

	Tracer trace.Tracer
}

// Client is a low-level way to interact with a Metio compliant endpoint.
type Client struct {
	client *nats.Conn
	// The prefix that is applied to names that are used when publishing messages
	prefix string

	// The tracer used to tracer the client
	tracer trace.Tracer
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

	tracer := config.Tracer
	if tracer == nil {
		// If no tracer is provided, use a noop tracer
		// to avoid always checking for nil
		tracer = trace.NewNoopTracerProvider().Tracer("noop")
	}
	return &Client{client: client, prefix: prefix, tracer: tracer}, nil
}

// Close closes the connection to the server.
func (c *Client) Close() {
	c.client.Close()
}

// Publish sends the given data to the server.
func (c *Client) Publish(ctx context.Context, subject string, data []types.Event) error {
	ctx, span := c.tracer.Start(ctx, "client.Client.Publish")
	defer span.End()
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

// GetTracer returns the tracer used by the Client
func (c *Client) GetTracer() trace.Tracer {
	return c.tracer
}
