package main

import (
	"github.com/Chris2010l/datapack-cli/src"
)

func main() {
	user_inputs := src.GetUserInputs()
	data_structure, err := src.GetDataStructure()
	if err != nil {
		return
	}

	data_structure, err = src.ReplaceYamlWithUserInputs(user_inputs)
	src.GenerateFolderStructure(user_inputs, data_structure)
}
