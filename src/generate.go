package src

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

func GenerateFolderStructure(user_options map[string]string, yaml_bytes []byte) error {
	config, err := ParseYaml(yaml_bytes)
	if err != nil {
		fmt.Printf("%s", err)
		return err
	}

	current_path := "."

	for _, step := range config.Steps {
		if step.Goto != "" {
			resolved_goto := step.Goto
			for placeholder, value := range user_options {
				resolved_goto = string(bytes.ReplaceAll([]byte(resolved_goto), []byte(placeholder), []byte(value)))
			}
			current_path = filepath.Join(current_path, resolved_goto)
			continue
		}

		for _, item := range step.Create {
			resolved_path := item.Path
			for placeholder, value := range user_options {
				resolved_path = string(bytes.ReplaceAll([]byte(resolved_path), []byte(placeholder), []byte(value)))
			}
			full_path := filepath.Join(current_path, resolved_path)

			// 1. Wenn es ein Ordner ist: Anlegen und fertig!
			if item.Type == "folder" {
				fmt.Printf("Creating Folder: %s\n", full_path)
				if err := os.MkdirAll(full_path, 0755); err != nil {
					fmt.Printf("%s", err)
					return err
				}
				continue // Wichtig: Überspringt den Dateicode für diesen Ordner
			}

			// 2. Wenn es eine Datei ist: Einlesen, Platzhalter ersetzen und schreiben
			if item.Type == "file" {
				// Stelle sicher, dass der Zielordner existiert
				if err := os.MkdirAll(filepath.Dir(full_path), 0755); err != nil {
					fmt.Printf("%s", err)
					return err
				}

				asset_path := filepath.Join("assets", item.Import)
				content, err := yaml_files.ReadFile(asset_path)
				if err != nil {
					// Fallback
					content, err = yaml_files.ReadFile(filepath.Join("assets", filepath.Base(item.Import)))
					if err != nil {
						fmt.Printf("Warning: Asset: %s could not be found: %s\n", item.Import, err)
						content = []byte("")
					}
				}

				for placeholder, value := range user_options {
					if placeholder == "&FORMAT" {
						old_bytes := []byte(`"` + placeholder + `"`)
						content = bytes.ReplaceAll(content, old_bytes, []byte(value))
						continue
					}
					content = bytes.ReplaceAll(content, []byte(placeholder), []byte(value))
				}

				fmt.Printf("Creating File: %s\n", full_path)
				if err := os.WriteFile(full_path, content, 0644); err != nil {
					fmt.Printf("Error: %s\n", err)
					return err
				}
			}
		}
	}
	return nil
}
