package unmarshaller

import (
	"encoding/json"
)

func DecodeJSON(body []byte) map[string]interface{} {
	var result map[string]interface{}
	json.Unmarshal(body, result)
	return result
}
