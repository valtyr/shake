package main

import "testing"

import (
	"github.com/stretchr/testify/assert"
	unmarshaller "github.com/valtyr/shake/pkg/unmarshaller"
)

func TestUnmarshaller(t *testing.T) {
	rawJSON := []byte("{\"type\": \"honk\"}")
	result := unmarshaller.DecodeJSON(rawJSON)
	assert.NotNil(t, result["type"], "Event type should not be nil")
}
