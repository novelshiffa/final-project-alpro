package types

import "fmt"

type Menu struct {
	Items                [NMAX]Text
	DefaultSelectedColor string
	Length               int
}

func (m *Menu) SetSelected(index int) {
	for i := 0; i < 3; i++ {
		if i == index {
			(*m).Items[i].SetColor((*m).DefaultSelectedColor)
		} else {
			(*m).Items[i].SetColor("white")
		}
	}
}

func (m *Menu) ShowAll() {
	for i := 0; i < m.Length; i++ {
		fmt.Println(m.Items[i].Colored)
	}
}
