package handler

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func InputTime(prompt string, attr *time.Time, required bool, dateOnly bool) {
	var yyyymmdd, hhmmss, datetimeStr string
	var promptText types.Text = types.NewText(prompt)

	// Layout for parsing the datetime string
	var layout string
	var format string

	if !dateOnly {
		layout = "2006-01-02 15:04:05"
		format = "yyyy-mm-dd hh:mm:ss"
	} else {
		layout = "2006-01-02"
		format = "yyyy-mm-dd"
	}

	var errorText = types.NewText("Invalid type of input, expects " + format + " string format. Try again.")
	errorText.SetColor("red")

	for {
		fmt.Print(promptText.Colored)
		fmt.Scanln(&yyyymmdd, &hhmmss)

		if (yyyymmdd == "" || hhmmss == "") && !required {
			return
		}

		if !dateOnly {
			datetimeStr = yyyymmdd + " " + hhmmss
		} else {
			datetimeStr = yyyymmdd
		}

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

func InputColumnName(_struct_ string, prmpt string, columnPtr *string) {
	var prompt types.Text = types.NewText(prmpt + " ")
	prompt.SetColor("white")

	var rightArrowText types.Text = types.NewText("[→] ")
	rightArrowText.SetColor("blue")

	var zeroToCancelText types.Text = types.NewText("(0 to cancel) ")
	zeroToCancelText.SetColor("red")

	var invalidInputErrText types.Text = types.NewText("Undefined column name. Try again.")
	invalidInputErrText.SetColor("red")

	var stopInput bool = false
	var temp string

	for !stopInput {
		fmt.Print(rightArrowText.Colored + prompt.Colored + zeroToCancelText.Colored)
		fmt.Scanln(&temp)

		var validInput bool

		if _struct_ == "items" || _struct_ == "item" {
			var items types.Items
			validInput = items.IsColumn(temp) || temp == "0"
		} else if _struct_ == "transactions" || _struct_ == "transaction" {
			var transactions types.Transactions
			validInput = transactions.IsColumn(temp) || temp == "0"
		}

		if validInput {
			stopInput = true
		} else {
			fmt.Println(invalidInputErrText.Colored)
		}
	}

	*columnPtr = temp
}

func InputlnString(ptr *string) {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	*ptr = strings.TrimSpace(input)
}
