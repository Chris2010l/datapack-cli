package util

import (
	"os"
	"path/filepath"
	"runtime"
)

func GetDataStructure() ([]byte, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, os.ErrNotExist
	}
	base_dir := filepath.Dir(filename)
	absoulte_path := filepath.Join(base_dir, "../datapack-structure.json")

	bytes, err := os.ReadFile(absoulte_path)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
