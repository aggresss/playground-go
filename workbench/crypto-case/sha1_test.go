package sha1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSha1(t *testing.T) {
	testCases := map[string]struct {
		signature string
		secret    string
		time      string
		nonce     string
		expected  bool
	}{
		"test1": {
			signature: "2b6bd834bb5c2a35166483c64e654749dd154fe1",
			secret:    "secret",
			time:      "1610698357000",
			nonce:     "1234567890",
			expected:  true,
		},
	}
	for n, c := range testCases {
		t.Run(n, func(t *testing.T) {
			assert.Equal(t, c.expected, VerifyMessageContent(c.signature, c.secret, c.time, c.nonce))
		})
	}
}
