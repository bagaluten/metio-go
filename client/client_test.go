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

package client_test

import (
	"context"
	"testing"

	"github.com/bagaluten/metio-go/client"
	"github.com/bagaluten/metio-go/types"
	"github.com/stretchr/testify/require"
)

func TestSerialization(t *testing.T) {
	ctx := context.Background()
	client, err := client.NewClient(client.Config{Host: "localhost:4222", Prefix: nil})
	require.NoError(t, err)

	defer client.Close()

	event := types.Event{
		EventID:   "123",
		ContextID: nil,
		EventType: types.MustParseEventType("group/name/version"),
		Payload: types.Payload{
			"key": "value",
		},
		Timestamp: types.TimeNow(),
	}

	err = client.Publish(ctx, "subject", []types.Event{event})
	require.NoError(t, err)

	client.Close()
}
