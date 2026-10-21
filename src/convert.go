package src

import (
	"bytes"
	"embed"
	"fmt"
	"io/fs"
	"path/filepath"
)

//go:embed assets/*
var yaml_files embed.FS

func loadAllYamlFiles() ([]byte, error) {
	var snippets [][]byte
	err := fs.WalkDir(yaml_files, "assets", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("%s", err)
			return err
		}

		if d.IsDir() {
			return nil
		}

		ext := filepath.Ext(path)
		if ext != ".yml" && ext != ".yaml" {
			return nil
		}

		content, err := yaml_files.ReadFile(path)
		if err != nil {
			fmt.Printf("%s", err)
			return err
		}

		snippets = append(snippets, content)

		return nil
	})

	if err != nil {
		fmt.Printf("%s", err)
		return nil, err
	}

	finalCombined := combineYaml(snippets...)

	return finalCombined, nil

}

func combineYaml(snippets ...[]byte) []byte {
	var combined []byte
	for _, snippet := range snippets {
		combined = append(combined, snippet...)
		combined = append(combined, '\n')
	}
	fmt.Printf("%s", combined)
	return combined
}

func ReplaceYamlWithUserInputs(user_options map[string]string) ([]byte, error) {
	default_yaml, err := loadAllYamlFiles()
	if err != nil {
		fmt.Printf("%s", err)
		return nil, err
	}
	for placeholder, value := range user_options {
		if placeholder == "&FORMAT" {
			oldBytes := []byte(`"` + placeholder + `"`)
			newBytes := []byte(value)
			default_yaml = bytes.ReplaceAll(default_yaml, oldBytes, newBytes)
			continue
		}

		default_yaml = bytes.ReplaceAll(default_yaml, []byte(placeholder), []byte(value))
	}

	return default_yaml, nil
}
