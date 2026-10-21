package src

import (
	"bytes"
	"testing"
)

func TestReplaceYamlWithUserInputs(t *testing.T) {
	mockUserOptions := map[string]string{
		"&DATAPACK":    "My cool datapack",
		"&FORMAT":      "117",
		"&DESCRIPTION": "Anoy my friends!!!",
		"&NAMESPACE":   "dp",
	}

	resultJson, err := ReplaceYamlWithUserInputs(mockUserOptions)
	if err != nil {
		t.Fatalf("Err: %s", err)
	}

	if !bytes.Contains(resultJson, []byte(`117`)) {
		t.Error("Error &FORMAT wasnt replaced")
	}

	if !bytes.Contains(resultJson, []byte("My cool datapack")) {
		t.Error("Name was skipped")
	}

	if !bytes.Contains(resultJson, []byte("Anoy my friends!!!")) {
		t.Error("Description wasnt replaced or was skipped")
	}

	if !bytes.Contains(resultJson, []byte("dp")) {
		t.Error("Namespace was replaced")
	}

	t.Log(string(resultJson))

}
