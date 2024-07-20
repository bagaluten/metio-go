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
	"fmt"
	"strings"
)

// Config is the configuration for the client.
type PartialError struct {
	errors []struct {
		Error error
		index int
	}
}

// GetFailedIndices returns the indices of the messages that failed to be delivered.
func (e PartialError) GetFailedIndices() []uint {
	indices := []uint{}
	for _, err := range e.errors {
		indices = append(indices, uint(err.index))
	}
	return indices
}

// Error returns a string representation of the error.
func (e PartialError) Error() string {
	stringBuilder := strings.Builder{}
	stringBuilder.WriteString(fmt.Sprintf("The delivery of %d messages failed\n", len(e.errors)))
	for _, err := range e.errors {
		stringBuilder.WriteString(fmt.Sprintf("\tMessage %d failed with error: %s", err.index, err.Error))
	}
	return stringBuilder.String()
}

// addError adds an error to the PartialError.
// it is not exported because it is only used internally.
func (e *PartialError) addError(err error, index int) {
	e.errors = append(e.errors, struct {
		Error error
		index int
	}{Error: err, index: index})
}

// isEmpty returns true if the PartialError is empty.
// it is not exported because it is only used internally.
// if you want to return empty return nil instead.
func (e *PartialError) isEmpty() bool {
	return len(e.errors) == 0
}
