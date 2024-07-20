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

package streams

import (
	"github.com/bagaluten/metio-go/client"
	"github.com/bagaluten/metio-go/types"
)

type Stream struct {
	// The name of the stream
	Name string

	// The client used to interact with the stream
	client *client.Client
}

func NewStream(name string, client *client.Client) *Stream {
	return &Stream{Name: name, client: client}
}

// Publish sends the given data to the server.
func (s Stream) Publish(data []types.Event) error {
	return s.client.Publish(s.Name, data)
}