package event

import (
	"encoding/json"
	"errors"
)

type Event struct {
	Type   string
	Tags   map[string]interface{}
	Fields map[string]interface{}
}

func DecodeJSON(body []byte) (Event, error) {
	var event Event

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	eventType, ok := result["type"].(string)
	if !ok {
		return event, errors.New("no type provided for event")
	}
	event.Type = eventType

	if result["tags"] != nil {
		tags, ok := result["tags"].(map[string]interface{})
		if !ok {
			return event, errors.New("the field event.tags must be an object")
		}
		event.Tags = tags
	}
	if result["fields"] != nil {
		fields, ok := result["fields"].(map[string]interface{})
		if !ok {
			return event, errors.New("the field event.fields must be an object")
		}
		event.Fields = fields
	}

	return event, nil
}
