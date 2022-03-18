package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeduplicateStringSlice(t *testing.T) {
	testCases := map[string]struct {
		input    []string
		expected []string
	}{
		"nil": {
			input:    nil,
			expected: nil,
		},
		"blank": {
			input:    []string{},
			expected: []string{},
		},
		"case 1": {
			input:    []string{"alpha", "alpha", "beta", "gamma", "gamma"},
			expected: []string{"alpha", "beta", "gamma"},
		},
		"case 2": {
			input:    []string{"iota", "kappa", "lambda", "mu", "nu", "nu", "xi", "omicron", "pai"},
			expected: []string{"iota", "kappa", "lambda", "mu", "nu", "xi", "omicron", "pai"},
		},
	}
	for n, c := range testCases {
		t.Run(n, func(t *testing.T) {
			assert.Equal(t, c.expected, DeduplicateStringSlice(c.input))
		})
	}
}
