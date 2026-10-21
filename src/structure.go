package src

import (
	"fmt"
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
	absoulte_path := filepath.Join(base_dir, "./assets/_main.yml")

	bytes, err := os.ReadFile(absoulte_path)
	if err != nil {
		fmt.Printf("%s", err)
		return nil, err
	}

	return bytes, nil
}
