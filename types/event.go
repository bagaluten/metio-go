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

package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

// EventType is a struct that represents the type of an event
type EventType struct {
	Group   string `json:"group"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

func (e *EventType) MarshalJSON() ([]byte, error) {
	// use String() to represent json value
	return json.Marshal(e.String())
}

func (e *EventType) UnmarshalJSON(data []byte) error {
	// parse json value to EventType
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	eventType, err := FromString(s)
	if err != nil {
		return fmt.Errorf("failed to parse event type: %w", err)
	}

	e.Group = eventType.Group
	e.Name = eventType.Name
	e.Version = eventType.Version
	return nil
}

// Implement String() method for EventType
func (e *EventType) String() string {
	return e.Group + "/" + e.Name + "/" + e.Version
}

func FromString(s string) (EventType, error) {
	// parse string to EventType
	split := strings.Split(s, "/")
	if len(split) != 3 {
		return EventType{}, errors.New("invalid event type")
	}

	return EventType{
		Group:   split[0],
		Name:    split[1],
		Version: split[2],
	}, nil
}

// Event is one of the main Metio types. It represents an event that happend at a specific point in time.
type Event struct {
	// EventID is a unique identifier for the event
	EventID string `json:"eventID"`

	// EventType is the type of the event
	EventType EventType `json:"eventType"`

	// ContextID is a unique identifier for the context in which the event was created.
	ContextID *string `json:"contextID,omitempty"`

	// Timestamp is the time at which the event was created in UTC.
	Timestamp time.Time `json:"timestamp"`

	// Payload is a map of key-value pairs that contain the event data.
	Payload map[string]string `json:"payload"`
}
