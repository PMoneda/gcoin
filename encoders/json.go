package encoders

import "encoding/json"

func ToJSON(data interface{}) string {
	dt, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(dt)
}
