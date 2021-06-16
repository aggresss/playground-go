package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type rho struct {
	Epsilon int    `json:"epsilon"`
	Zeta    string `json:"zeta"`
	Eta     uint8  `json:"eta"`
	Iota    uint32
}

func TestTransJsonTag(t *testing.T) {
	testCases := map[string]struct {
		input    interface{}
		expected string
	}{
		"blank": {
			input:    &rho{},
			expected: "",
		},
		"full": {
			input: &rho{
				Epsilon: 8,
				Zeta:    "bar",
				Eta:     99,
			},
			expected: "epsilon=8;zeta=bar;eta=99",
		},
	}

	for n, c := range testCases {
		t.Run(n, func(t *testing.T) {
			assert.Equal(t, c.expected, transJsonTag(c.input))
		})
	}
}
