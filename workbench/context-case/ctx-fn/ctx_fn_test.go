package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCtxFn(t *testing.T) {
	ctx := context.Background()
	newCtx, cancel := context.WithCancel(ctx)
	assert.NotEqual(t, ctx, newCtx)
	defer cancel()
}
