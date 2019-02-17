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

func processField(field interface{}) (interface{}, error) {
	switch field.(type) {
	case float64:
		return field, nil
	case string:
		return field, nil
	default:
		return nil, errors.New("nested objects not permitted")
	}
}

func DecodeJSON(body []byte) (Event, error) {
	var event Event

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	tagLen := len(result) - 1

	eventType, ok := result["type"].(string)
	if !ok {
		return event, errors.New("no type provided for event")
	}
	event.Type = eventType

	if result["data"] != nil {
		fields, ok := result["data"].(map[string]interface{})
		if !ok {
			return event, errors.New("the field event.data must be an object")
		}

		event.Fields = make(map[string]interface{}, len(fields))
		for k := range fields {
			processedField, err := processField(fields[k])
			if err != nil {
				return event, err
			}
			event.Fields[k] = processedField
		}

		event.Fields = fields
		tagLen--
	}

	event.Tags = make(map[string]interface{}, tagLen)
	for k := range result {
		switch k {
		case "type":
			fallthrough
		case "data":
			continue
		default:
			processedField, err := processField(result[k])
			if err != nil {
				return event, err
			}
			event.Tags[k] = processedField
		}
	}

	return event, nil
}
