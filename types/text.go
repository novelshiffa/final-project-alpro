package types

type Text struct {
	Value   string
	Color   string
	Colored string
}

func NewText(Value string) Text {
	var text = Text{Value: Value, Color: "white", Colored: ""}
	text.SetColor("white")

	return text
}

func (t *Text) SetColor(Color string) {
	code := "\033["
	reset := code + "0m"

	switch Color {
	case "red":
		code += "31m"
	case "green":
		code += "32m"
	case "yellow":
		code += "33m"
	case "blue":
		code += "34m"
	case "white":
		code += "37m"
	default:
		code = ""
		reset = ""
	}

	(*t).Colored = code + (*t).Value + reset

	//fmt.Println((*t).Colored)
}
