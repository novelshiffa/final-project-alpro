package handler

import (
	"fmt"
	"strconv"
	"time"

	"github.com/novelshiffa/final-project-alpro/types"
)

func InputItem(prompt string, attr *types.Item, required bool, itemsRef *types.Items) {
	var temp int

	var promptText = types.NewText(prompt)
	var errorText = types.NewText("Item not found. Try again")
	errorText.SetColor("red")

	var index int

	for {
		fmt.Print(promptText.Colored)
		fmt.Scanln(&temp)

		if temp == 0 && !required {
			return
		}

		index = itemsRef.FindById(temp)

		if index != -1 {
			*attr = itemsRef.Items[index]

			return
		}

		fmt.Println(errorText.Colored)
	}
}

func InputTransactionType(prompt string, attr *string, required bool) {
	var temp string

	var promptText types.Text = types.NewText(prompt)
	var errorText = types.NewText("Invalid input, expects either incoming or outgoing. Try again.")
	errorText.SetColor("red")

	for {
		fmt.Print(promptText.Colored)
		fmt.Scanln(&temp)

		if temp == "" && !required {
			return
		}

		if temp == "incoming" || temp == "outgoing" {
			*attr = temp
			return
		}

		fmt.Println(errorText.Colored)

		temp = ""
	}
}

func InputTime(prompt string, attr *time.Time, required bool) {
	var yyyymmdd, hhmmss, datetimeStr string
	var promptText types.Text = types.NewText(prompt)

	// Layout for parsing the datetime string
	var layout string = "2006-01-02 15:04:05"

	var errorText = types.NewText("Invalid type of input, expects yyyy-mm-dd hh:mm:ss string format. Try again.")
	errorText.SetColor("red")

	for {
		fmt.Print(promptText.Colored)
		fmt.Scanln(&yyyymmdd, &hhmmss)

		if (yyyymmdd == "" || hhmmss == "") && !required {
			return
		}

		datetimeStr = yyyymmdd + " " + hhmmss

		// Parse the combined string into a time.Time object
		datetime, err := time.Parse(layout, datetimeStr)

		if err == nil {
			*attr = datetime
			return
		}

		fmt.Println(errorText.Colored)

		yyyymmdd, hhmmss = "", ""
	}
}

func InputInteger(prompt string, attr *int, required bool) {
	var temp string
	var err error
	var val int

	var promptText types.Text = types.NewText(prompt)
	var errorText = types.NewText("Invalid type of input, expects an integer. Try again.")
	errorText.SetColor("red")

	for {
		fmt.Print(promptText.Colored)
		fmt.Scanln(&temp)

		if temp == "" && !required {
			return
		}

		val, err = strconv.Atoi(temp)

		if err == nil {
			*attr = val
			return
		}

		fmt.Println(errorText.Colored)

		temp = ""
	}
}
