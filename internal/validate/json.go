package validate

import (
	"encoding/json"
)

// JSON validates a JSON string, it returns a boolean indicating whether
// the JSON is valid or not.
func JSON(jsonStr string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(jsonStr), &js) == nil
}
