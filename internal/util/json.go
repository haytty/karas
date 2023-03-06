package util

import "encoding/json"

func PrettyJSON(v any) ([]byte, error) {
	return json.MarshalIndent(v, "", "    ")
}
