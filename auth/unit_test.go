package auth

import (
	"github.com/stretchr/testify/assert"
	"hexnet/api/common"
	"testing"
)

func TestNewAuthMiddleware(t *testing.T) {
	asserts := assert.New(t)
	common.LoadConfig("../.env")

	defer func() {
		if rec := recover(); rec != nil {
			asserts.True(false, "Middleware could not be initialized: panic caught")
		}
	}()
	m := NewAuthMiddleware()
	asserts.NotNil(m, "Middleware pointer is nil")
}
