package new

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewElemGamma(t *testing.T) {
	rho := WithFoo(true)
	sigma := WithBar(99)
	tau := WithZoo("test")
	gamma, err := NewElemGamma(rho, sigma, tau)
	assert.Nil(t, err)
	assert.True(t, gamma.Foo)
	assert.Equal(t, 99, gamma.Bar)
	assert.Equal(t, "test", gamma.Zoo)
}
