package main

import (
	"github.com/Chris2010l/datapack-cli/util"
)

func main() {
	user_inputs := util.GetUserInputs()
	data_structure, err := util.GetDataStructure()
	if err != nil {
		return
	}

	data_structure, err = util.ReplaceJsonToUserJson(user_inputs, data_structure)

	util.GenerateFolderStructure(user_inputs, data_structure)
}
