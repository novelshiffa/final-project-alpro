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

// SetSelected mengatur warna teks untuk item menu tertentu berdasarkan indeks yang dipilih, dan mengubah warna item yang lain menjadi warna default.
//
// Parameter:
// - index (int): Indeks item menu yang dipilih untuk diubah warnanya.
func (m *Menu) SetSelected(index int) {
	for i := 0; i < m.Length; i++ {
		if i == index {
			(*m).Items[i].SetColor((*m).DefaultSelectedColor) // Mengatur warna item yang dipilih menjadi warna default yang ditentukan
		} else {
			(*m).Items[i].SetColor("white") // Mengatur warna item yang lain menjadi putih
		}
	}
}

// ShowAll menampilkan semua item menu dalam Menu beserta warna teks yang sudah diatur ke dalam Colored.
// Setiap item menu ditampilkan di baris baru.
func (m *Menu) ShowAll() {
	for i := 0; i < m.Length; i++ {
		fmt.Println(m.Items[i].Colored) // Menampilkan teks berwarna untuk setiap item menu
	}
}

// Listen mengaktifkan mode mendengarkan keyboard untuk interaksi menu interaktif.
// Fungsi ini memungkinkan navigasi atas-bawah menggunakan tombol panah untuk memilih item menu,
// dan menanggapi tombol "Enter" untuk menghentikan loop mendengarkan dan menjalankan aksi yang diberikan.
//
// Parameter:
// - selector (*int): Pointer ke variabel integer yang menunjukkan indeks item menu yang dipilih saat ini.
// - stopLoop (*bool): Pointer ke variabel boolean yang menentukan apakah loop mendengarkan harus dihentikan.
// - clear (*bool): Pointer ke variabel boolean yang menunjukkan apakah terminal harus dihapus sebelum menampilkan menu.
// - action (func()): Fungsi yang akan dijalankan setelah pengguna memilih item menu dan menekan "Enter".
// - prevAction (func()): Fungsi yang akan dijalankan sebelum menampilkan menu, biasanya digunakan untuk mempersiapkan kondisi sebelumnya.
func (m *Menu) Listen(selector *int, stopLoop *bool, clear *bool, action func(), prevAction func()) {
	var stopListening bool = false

	for !stopListening {
		// Membersihkan terminal jika clear diaktifkan
		if *clear {
			utils.ClearTerminal()
		}

		// Menjalankan aksi sebelum menampilkan menu
		prevAction()

		// Menampilkan semua item menu dengan warna yang sudah diatur
		(*m).ShowAll()

		// Mendengarkan input keyboard menggunakan library atomic keyboard
		keyboard.Listen(func(key keys.Key) (stop bool, err error) {
			// Mengatur selector untuk navigasi atas-bawah
			if key.Code == keys.Up && *selector > 0 {
				*selector--
				(*m).SetSelected(*selector)
			} else if key.Code == keys.Down && *selector < (*m).Length-1 {
				*selector++
				(*m).SetSelected(*selector)
			}

			// Menghentikan loop jika tombol "Enter" ditekan
			if key.Code == keys.Enter {
				*stopLoop = true
			}

			return true, nil
		})

		// Menjalankan aksi setelah pengguna memilih item menu
		if *stopLoop {
			action()

			// Memastikan untuk menghentikan loop jika stopLoop tetap true setelah aksi
			if *stopLoop {
				stopListening = true
			}
		}
	}
}
