package handler

import (
	"fmt"
	"strconv"

	"github.com/novelshiffa/final-project-alpro/types"
)

func InputInteger(prompt string, value *int, required bool) {
	var temp string
	var err error
	var val int

	for {
		fmt.Print(prompt)
		fmt.Scanln(&temp)

		if temp == "" && !required {
			return
		}

		val, err = strconv.Atoi(temp)

		if err == nil {
			*value = val
			return
		}

		var errorText = types.NewText("Invalid type of input, expects an integer. Try again.")
		errorText.SetColor("red")
		fmt.Println(errorText.Colored)

		temp = ""
	}
}
