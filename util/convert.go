package util

import "bytes"

func ReplaceJsonToUserJson(user_options map[string]string, default_json []byte) ([]byte, error) {
	for placeholder, value := range user_options {
		if placeholder == "&NAME" {
			continue
		}

		if placeholder == "&FORMAT" {
			oldBytes := []byte(`"` + placeholder + `"`)
			newBytes := []byte(value)
			default_json = bytes.ReplaceAll(default_json, oldBytes, newBytes)
			continue
		}

		default_json = bytes.ReplaceAll(default_json, []byte(placeholder), []byte(value))
	}

	return default_json, nil
}
