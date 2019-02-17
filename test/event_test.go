package main

import "testing"

import (
	"github.com/stretchr/testify/assert"
	event "github.com/valtyr/shake/pkg/event"
)

func TestEventParsing(t *testing.T) {
	rawJSON := []byte(`{"type": "app_opened", "user_id": "123", "data": {"active_screen": "home"}}`)
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

func TestEventParsingDataError(t *testing.T) {
	rawJSON := []byte(`{"type": "app_opened", "data": 4}`)
	_, err := event.DecodeJSON(rawJSON)
	assert.Error(t, err)
}

func TestEventParsingTags(t *testing.T) {
	rawJSON := []byte(`{"type": "app_opened", "data": {"test": "abc"}, "user_id": "abc123", "number": 2}`)
	result, _ := event.DecodeJSON(rawJSON)
	assert.Equal(t, result.Tags["user_id"], "abc123")
	assert.Equal(t, result.Tags["number"], float64(2))
	assert.Equal(t, result.Fields["test"], "abc")
}

func TestEventParsingFieldsError(t *testing.T) {
	rawJSON := []byte(`{"type": "app_opened", "data": {"test": {"abc": 123}}, "user_id": "abc123"}`)
	_, err := event.DecodeJSON(rawJSON)
	assert.Error(t, err)
}

func TestEventParsingTagsError(t *testing.T) {
	rawJSON := []byte(`{"type": "app_opened", "data": {"test": 123}, "user_id": {"test": {"abc": 123}}}`)
	_, err := event.DecodeJSON(rawJSON)
	assert.Error(t, err)
}
