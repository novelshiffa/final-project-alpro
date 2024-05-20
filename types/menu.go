package types

import (
	"fmt"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/novelshiffa/final-project-alpro/utils"
)

type Menu struct {
	Items                [NMAX]Text
	DefaultSelectedColor string
	Length               int
}

func (m *Menu) SetSelected(index int) {
	for i := 0; i < m.Length; i++ {
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

func (m *Menu) Listen(selector *int, stopLoop *bool, clear *bool, action func(), prevAction func()) {
	prevAction()
	for {
		if *clear {
			utils.ClearTerminal()
		}

		(*m).ShowAll()

		keyboard.Listen(func(key keys.Key) (stop bool, err error) {
			if key.Code == keys.Up && *selector > 0 {
				*selector--
				(*m).SetSelected(*selector)
			} else if key.Code == keys.Down && *selector < (*m).Length-1 {
				*selector++
				(*m).SetSelected(*selector)
			}

			if key.Code == keys.Enter {
				*stopLoop = true
			}

			return true, nil
		})

		if *stopLoop {
			action()
		}
	}
}
