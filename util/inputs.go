package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserInputs() map[string]string {
	user_inputs := make(map[string]string)

	scanner := bufio.NewScanner(os.Stdin)

	ask := func(question string) string {
		fmt.Println(question)
		if scanner.Scan() {
			return strings.TrimSpace(scanner.Text())
		}
		return ""
	}

	user_inputs["&DATAPACK"] = ask("What name should be used for the datapack?")
	user_inputs["&NAMESPACE"] = ask("What namespace should be used for the datapack?")
	user_inputs["&FORMAT"] = ask("What datapack format should be used for the datapack?")
	user_inputs["&DESCRIPTION"] = ask("What datapack description should be used for the datapack?")

	return user_inputs
}
