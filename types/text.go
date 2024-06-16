package types

type Text struct {
	Value   string
	Color   string
	Colored string
}

// NewText membuat instance baru dari struct Text dengan nilai teks yang diberikan dan warna default "white".
// Fungsi ini juga memanggil SetColor("white") untuk menginisialisasi teks dengan warna putih.
//
// Parameter:
// - Value (string): Nilai teks yang ingin dimasukkan ke dalam struktur Text.
//
// Return:
// - Text: Instance baru dari struct Text dengan nilai dan warna default.
func NewText(Value string) Text {
	var text = Text{Value: Value, Color: "white", Colored: ""}
	text.SetColor("white")

	return text
}

// SetValue mengatur nilai teks baru untuk instance Text dan secara otomatis memperbarui teks berwarna.
//
// Parameter:
// - text (string): Nilai teks baru yang ingin diatur.
func (t *Text) SetValue(text string) {
	t.Value = text
	t.SetColor(t.Color) // Memanggil SetColor untuk memperbarui teks berwarna berdasarkan warna saat ini
}

// SetColor mengatur warna teks untuk instance Text berdasarkan nilai warna yang diberikan.
// Fungsi ini juga mengupdate properti Colored dengan teks berwarna menggunakan kode ANSI.
//
// Parameter:
// - Color (string): String yang menyatakan warna teks yang diinginkan. Nilai yang diterima adalah: "red", "green", "yellow", "blue", "white", "cyan", "magenta", "black".
func (t *Text) SetColor(Color string) {
	code := "\033["      // Kode ANSI untuk mengatur warna teks di terminal
	reset := code + "0m" // Kode ANSI untuk mereset warna ke default

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
	case "cyan":
		code += "36m"
	case "magenta":
		code += "35m"
	case "black":
		code += "30m"
	default:
		code = "" // Jika warna tidak valid, kode ANSI dikosongkan
		reset = ""
	}

	(*t).Color = Color
	(*t).Colored = code + (*t).Value + reset // Memformat teks dengan kode ANSI sesuai warna yang dipilih
}
