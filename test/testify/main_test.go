package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	ast := assert.New(t)
	name := "Bob"
	age := 20

	ast.Equal("Bob", name)
	ast.Equal(20, age)
}

func TestCase2(t *testing.T) {
	ast := require.New(t)
	name := "Bob"
	age := 20

	ast.Equal("Bob", name)
	ast.Equal(20, age)
}
