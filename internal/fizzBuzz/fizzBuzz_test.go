package fizzBuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fizzBuzzValidation(t *testing.T) {
	cases := []struct {
		name     string
		value    int
		expected string
	}{
		{
			name:     "FizzBuzz",
			value:    15,
			expected: "FizzBuzz",
		},
		{
			name:     "Fizz",
			value:    3,
			expected: "Fizz",
		},
		{
			name:     "Buzz",
			value:    5,
			expected: "Buzz",
		},
		{
			name:     "unknown",
			value:    1,
			expected: "1",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := verifyFizzBuzzInternal(tc.value)
			assert.Equal(t, tc.expected, result)
		})
	}
}
