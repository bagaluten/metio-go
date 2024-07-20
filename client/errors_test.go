package client

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestErrors(t *testing.T) {
	t.Run("TestPartialError", func(t *testing.T) {
		partialError := PartialError{}
		partialError.addError(errors.New("error"), 1)
		partialError.addError(errors.New("error"), 2)
		partialError.addError(errors.New("error"), 5)

		require.Equal(t, []uint{1, 2, 5}, partialError.GetFailedIndices())
		require.Equal(t, "The delivery of 3 messages failed\n\tMessage 1 failed with error: error\tMessage 2 failed with error: error\tMessage 5 failed with error: error", partialError.Error())

	})
}
