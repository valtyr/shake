package main

import "testing"

import (
	"github.com/stretchr/testify/assert"
	event "github.com/valtyr/shake/pkg/event"
)

func TestEventParsing(t *testing.T) {
	rawJSON := []byte(`{"type": "app_opened", "fields": {"active_screen": "home"}}`)
	result, _ := event.DecodeJSON(rawJSON)
	assert.Equal(t, result.Type, "app_opened")
	assert.Equal(t, result.Fields["active_screen"].(string), "home")
}
