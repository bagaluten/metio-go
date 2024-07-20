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

package types_test

import (
	"encoding/json"
	"testing"

	"github.com/bagaluten/metio-go/types"
	"github.com/stretchr/testify/require"
)

func TestSerialization(t *testing.T) {
	t.Run("TestSerializeEventType", func(t *testing.T) {
		eventType := types.EventType{
			Group:   "group",
			Name:    "name",
			Version: "version",
		}

		bytes, err := json.Marshal(eventType)
		require.NoError(t, err)

		require.JSONEq(t, `"group/name/version"`, string(bytes))

		var newEvent = types.EventType{}
		err = json.Unmarshal(bytes, &newEvent)
		require.NoError(t, err)
		require.Equal(t, eventType, newEvent)
	})

	t.Run("TestSerializeEvent", func(t *testing.T) {
		eventType, err := types.ParseEventType("group/name/version")
		require.NoError(t, err)

		event := types.Event{
			EventID:   "event-id",
			EventType: eventType,
			Payload: types.Payload{
				"key": "value",
			},
			Timestamp: types.TimeNow(),
			ContextID: nil,
		}

		bytes, err := json.Marshal(event)
		require.NoError(t, err)
		require.JSONEq(t, `{"eventID": "event-id", "eventType":"group/name/version","payload":{"key":"value"},"timestamp":"`+event.Timestamp.String()+`"}`, string(bytes))

		var newEvent = types.Event{}
		err = json.Unmarshal(bytes, &newEvent)
		require.NoError(t, err)

		require.Equal(t, event, newEvent)
	})

}
