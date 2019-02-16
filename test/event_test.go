package main

import "testing"

import (
	"github.com/stretchr/testify/assert"
	event "github.com/valtyr/shake/pkg/event"
)

func TestEventParsing(t *testing.T) {
	rawJSON := []byte(`{"type": "app_opened", "fields": {"active_screen": "home"}}`)
	result, err := event.DecodeJSON(rawJSON)
	assert.NoError(t, err)
	assert.Equal(t, result.Type, "app_opened")
	assert.Equal(t, result.Fields["active_screen"].(string), "home")
}

func TestEventParsingNoTypeError(t *testing.T) {
	rawJSON := []byte(`{"fields": {"active_screen": "home"}}`)
	_, err := event.DecodeJSON(rawJSON)
	assert.Error(t, err)
}

func TestEventParsingFieldsError(t *testing.T) {
	rawJSON := []byte(`{"type": "app_opened", "fields": 4}`)
	_, err := event.DecodeJSON(rawJSON)
	assert.Error(t, err)
}
