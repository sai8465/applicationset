package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombineStringMaps(t *testing.T) {
	testCases := []struct {
		name        string
		left        map[string]string
		right       map[string]string
		expected    map[string]string
		expectedErr error
	}{
		{
			name:        "combines the maps",
			left:        map[string]string{"foo": "bar"},
			right:       map[string]string{"a": "b"},
			expected:    map[string]string{"a": "b", "foo": "bar"},
			expectedErr: nil,
		},
		{
			name:        "fails if keys are the same but value isn't",
			left:        map[string]string{"foo": "bar", "a": "fail"},
			right:       map[string]string{"a": "b", "c": "d"},
			expected:    map[string]string{"a": "b", "foo": "bar"},
			expectedErr: errors.New("found duplicate key a with different value, a: fail ,b: b"),
		},
		{
			name:        "pass if keys & values are the same",
			left:        map[string]string{"foo": "bar", "a": "b"},
			right:       map[string]string{"a": "b", "c": "d"},
			expected:    map[string]string{"a": "b", "c": "d", "foo": "bar"},
			expectedErr: nil,
		},
	}

	for _, testCase := range testCases {
		testCaseCopy := testCase

		t.Run(testCaseCopy.name, func(t *testing.T) {
			t.Parallel()

			got, err := CombineStringMaps(testCaseCopy.left, testCaseCopy.right)

			if testCaseCopy.expectedErr != nil {
				assert.EqualError(t, err, testCaseCopy.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCaseCopy.expected, got)
			}

		})
	}
}
