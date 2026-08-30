package util

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Node struct {
	IsFolder bool            `json:"isFolder"`
	Content  json.RawMessage `json:"content"`
}

func GenerateFolderStructure(userOptions map[string]string, jsonBytes []byte) error {

	var root map[string]Node
	if err := json.Unmarshal(jsonBytes, &root); err != nil {
		return err
	}

	outputDir := "."

	return generateLevel(outputDir, root)
}

func generateLevel(currentPath string, nodes map[string]Node) error {
	if err := os.MkdirAll(currentPath, 0755); err != nil {
		return err
	}

	for name, node := range nodes {
		fullPath := filepath.Join(currentPath, name)

		if node.IsFolder {
			var subNodes map[string]Node
			if len(node.Content) > 0 && string(node.Content) != "{}" {
				if err := json.Unmarshal(node.Content, &subNodes); err != nil {
					return err
				}
				if err := generateLevel(fullPath, subNodes); err != nil {
					return err
				}
			} else {
				fmt.Printf("Creating Folder: %s\n", fullPath)
				if err := os.MkdirAll(fullPath, 0755); err != nil {
					return err
				}
			}
		} else {
			var rawObject interface{}
			if err := json.Unmarshal(node.Content, &rawObject); err != nil {
				return err
			}

			prettyJSON, err := json.MarshalIndent(rawObject, "", "	")
			if err != nil {
				return err
			}

			fmt.Printf("Creating File: %s\n", fullPath)
			if err := os.WriteFile(fullPath, prettyJSON, 0644); err != nil {
				return err
			}
		}
	}
	return nil
}
