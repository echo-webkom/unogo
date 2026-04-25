package test

import (
	"testing"

	"github.com/echo-webkom/unogo"
	"gotest.tools/v3/assert"
)

func TestHealth(t *testing.T) {
	client, err := unogo.NewClient(&unogo.Config{
		Url: "http://localhost:8000",
	})
	assert.NilError(t, err)

	resp, err := client.Health()
	assert.NilError(t, err)
	assert.Equal(t, resp, unogo.HealthResponse{Status: "ok"})
}
