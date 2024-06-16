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

// InputlnString meminta pengguna untuk memasukkan sebuah string dari stdin (standar input).
// Fungsi ini membaca input dari pengguna sampai menemukan karakter newline '\n'.
// Input yang dibaca kemudian dipotong spasi tambahan di awal dan akhir string menggunakan strings.TrimSpace.
// Hasil input yang sudah diproses kemudian disimpan ke dalam alamat yang ditunjuk oleh ptr.
//
// Parameter:
// - ptr (*string): Pointer ke string di mana hasil input akan disimpan.
func InputlnString(ptr *string) {
	// Membuat pembaca input dari stdin
	reader := bufio.NewReader(os.Stdin)

	// Membaca input dari pengguna sampai menemukan karakter newline '\n'
	input, _ := reader.ReadString('\n')

	// Menghapus spasi tambahan di awal dan akhir string
	*ptr = strings.TrimSpace(input)
}

// OldValueFormat mengembalikan teks yang diformat dengan tambahan "[!] Current value: " diikuti oleh nilai lama.
// Teks tersebut akan diberi warna biru.
//
// Parameter:
// - oldValue (string): Nilai lama yang ingin diformat dalam teks.
//
// Return:
// - string: Teks yang diformat dengan tambahan "[!] Current value: " diikuti oleh oldValue, berwarna biru.
func OldValueFormat(oldValue string) string {
	// Membuat objek teks baru dengan nilai "[!] Current value: " + oldValue dan memberi warna biru
	text := types.NewText("[!] Current value: " + oldValue + "\n")
	text.SetColor("blue")

	return text.Colored // Mengembalikan teks yang sudah diformat dan diberi warna
}

// RightArrowedPrompt mengembalikan teks yang diformat dengan tambahan "[→] " diikuti oleh teks prompt yang diberikan.
// Teks "[→] " dan teks prompt tersebut akan diberi warna biru.
//
// Parameter:
// - prompt (string): Teks prompt yang ingin ditampilkan setelah "[→] ".
//
// Return:
// - string: Teks yang diformat dengan tambahan "[→] " diikuti oleh prompt, semuanya berwarna biru.
func RightArrowedPrompt(prompt string) string {
	// Membuat objek teks baru dengan nilai "[→] " dan memberi warna biru
	var rightArrowText types.Text = types.NewText("[→] ")
	rightArrowText.SetColor("blue")

	// Menggabungkan "[→] " yang sudah diformat dengan prompt yang juga sudah diformat
	return rightArrowText.Colored + types.NewText(prompt).Colored
}

// InputItem meminta pengguna untuk memasukkan ID item dengan prompt tertentu, kemudian mencocokkan ID tersebut dengan item dalam struktur Items yang diberikan.
// Jika ID item ditemukan, atribut item yang diberikan (*attr) akan diisi dengan item yang cocok.
// Fungsi ini akan terus meminta input ulang jika ID tidak ditemukan, kecuali jika input 0 dan required=false.
//
// Parameter:
// - prompt (string): Teks prompt yang menunjukkan apa yang harus dimasukkan pengguna.
// - attr (*types.Item): Pointer ke objek Item tempat hasil pencarian akan disimpan.
// - required (bool): Jika true, input ID item diperlukan (tidak boleh 0). Jika false, input ID item bisa 0 untuk keluar dari fungsi.
// - itemsRef (*types.Items): Pointer ke struktur Items yang berisi item-item yang dicari berdasarkan ID.
func InputItem(prompt string, attr *types.Item, required bool, itemsRef *types.Items) {
	var temp int

	// Membuat teks prompt dan teks error untuk output
	var promptText = types.NewText(prompt)
	var errorText = types.NewText("Item not found. Try again.")
	errorText.SetColor("red")

	var index int

	for {
		// Menampilkan prompt untuk meminta input ID item
		fmt.Print(promptText.Colored)
		fmt.Scanln(&temp)

		// Jika input 0 dan tidak required, maka keluar dari fungsi
		if temp == 0 && !required {
			return
		}

		// Mencari index item dengan ID yang diinput dalam struktur Items
		index = itemsRef.FindById(temp)

		// Jika item ditemukan berdasarkan ID, mengisi *attr dengan item yang cocok dan keluar dari fungsi
		if index != -1 {
			*attr = itemsRef.Items[index]
			return
		}

		// Jika ID item tidak ditemukan, menampilkan pesan error
		fmt.Println(errorText.Colored)
	}
}

// InputTransactionType meminta pengguna untuk memasukkan tipe transaksi dengan prompt tertentu, kemudian memvalidasi input tersebut.
// Input yang valid adalah "incoming" atau "outgoing". Jika input valid, nilai atribut yang diberikan (*attr) akan diisi dengan input tersebut.
// Fungsi ini akan terus meminta input ulang jika input tidak valid, kecuali jika input kosong dan required=false, atau jika input sama dengan nilai exception.
//
// Parameter:
// - prompt (string): Teks prompt yang menunjukkan apa yang harus dimasukkan pengguna.
// - attr (*string): Pointer ke string tempat hasil input tipe transaksi akan disimpan.
// - required (bool): Jika true, input tipe transaksi diperlukan (tidak boleh kosong). Jika false, input kosong diizinkan untuk keluar dari fungsi.
// - exception (string): Nilai pengecualian yang akan menyebabkan fungsi keluar jika diinput.
func InputTransactionType(prompt string, attr *string, required bool, exception string) {
	var temp string

	// Membuat teks prompt dan teks error untuk output
	var promptText types.Text = types.NewText(prompt)
	var errorText = types.NewText("Invalid input, expects either incoming or outgoing. Try again.")
	errorText.SetColor("red")

	for {
		// Menampilkan prompt untuk meminta input tipe transaksi
		fmt.Print(promptText.Colored)
		fmt.Scanln(&temp)

		// Jika input kosong dan tidak required, atau jika input sama dengan exception, keluar dari fungsi
		if temp == "" && !required || temp == exception {
			return
		}

		// Memvalidasi input: harus "incoming" atau "outgoing"
		if temp == "incoming" || temp == "outgoing" {
			*attr = temp
			return
		}

		// Jika input tidak valid, menampilkan pesan error
		fmt.Println(errorText.Colored)

		// Mengosongkan temp untuk input ulang
		temp = ""
	}
}

// InputTime meminta pengguna untuk memasukkan waktu dengan prompt tertentu, kemudian memvalidasi input tersebut.
// Fungsi ini dapat meminta format tanggal dan waktu lengkap atau hanya tanggal, tergantung pada parameter dateOnly.
// Jika input valid, waktu yang di-parse akan disimpan dalam attr (pointer ke time.Time).
// Fungsi ini akan terus meminta input ulang jika format tidak sesuai atau jika input kosong dan required=false.
//
// Parameter:
// - prompt (string): Teks prompt yang menunjukkan apa yang harus dimasukkan pengguna.
// - attr (*time.Time): Pointer ke time.Time tempat hasil waktu yang di-parse akan disimpan.
// - required (bool): Jika true, input waktu diperlukan (tidak boleh kosong). Jika false, input kosong diizinkan untuk keluar dari fungsi.
// - dateOnly (bool): Jika true, hanya tanggal yang diminta dan di-parse. Jika false, tanggal dan waktu (hingga detik) diminta dan di-parse.
func InputTime(prompt string, attr *time.Time, required bool, dateOnly bool) {
	var yyyymmdd, hhmmss, datetimeStr string
	var promptText types.Text = types.NewText(prompt)

	// Layout untuk parsing string datetime
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
		// Meminta input dari pengguna
		fmt.Print(promptText.Colored)
		fmt.Scanln(&yyyymmdd, &hhmmss)

		// Jika input kosong dan tidak required, atau jika hanya tanggal dan tidak ada waktu yang diminta
		if (yyyymmdd == "" || hhmmss == "") && !required {
			return
		}

		// Menggabungkan tanggal dan waktu menjadi satu string datetime
		if !dateOnly {
			datetimeStr = yyyymmdd + " " + hhmmss
		} else {
			datetimeStr = yyyymmdd
		}

		// Parsing string datetime menjadi objek time.Time
		datetime, err := time.Parse(layout, datetimeStr)

		// Jika parsing berhasil, menyimpan hasil di attr dan keluar dari fungsi
		if err == nil {
			*attr = datetime
			return
		}

		// Jika parsing gagal, menampilkan pesan error
		fmt.Println(errorText.Colored)

		// Mengosongkan variabel untuk input ulang
		yyyymmdd, hhmmss = "", ""
	}
}

// InputInteger meminta pengguna untuk memasukkan angka bulat dengan prompt tertentu, kemudian memvalidasi input tersebut.
// Jika input valid, angka bulat yang di-parse akan disimpan dalam attr (pointer ke int).
// Fungsi ini akan terus meminta input ulang jika input tidak sesuai dengan format integer atau jika input kosong dan required=false.
//
// Parameter:
// - prompt (string): Teks prompt yang menunjukkan apa yang harus dimasukkan pengguna.
// - attr (*int): Pointer ke int tempat hasil integer yang di-parse akan disimpan.
// - required (bool): Jika true, input integer diperlukan (tidak boleh kosong). Jika false, input kosong diizinkan untuk keluar dari fungsi.
func InputInteger(prompt string, attr *int, required bool) {
	var temp string
	var err error
	var val int

	var promptText types.Text = types.NewText(prompt)
	var errorText = types.NewText("Invalid type of input, expects an integer. Try again.")
	errorText.SetColor("red")

	for {
		// Meminta input dari pengguna
		fmt.Print(promptText.Colored)
		fmt.Scanln(&temp)

		// Jika input kosong dan tidak required, keluar dari fungsi
		if temp == "" && !required {
			return
		}

		// Mengonversi string menjadi integer
		val, err = strconv.Atoi(temp)

		// Jika konversi berhasil, menyimpan hasil di attr dan keluar dari fungsi
		if err == nil {
			*attr = val
			return
		}

		// Jika konversi gagal, menampilkan pesan error
		fmt.Println(errorText.Colored)

		// Mengosongkan variabel untuk input ulang
		temp = ""
	}
}

// InputColumnName meminta pengguna untuk memasukkan nama kolom yang valid sesuai dengan jenis struktur yang diberikan.
// Fungsi ini memvalidasi input untuk menentukan apakah nama kolom yang dimasukkan valid untuk struktur yang spesifik (items atau transactions).
// Jika input valid, nama kolom akan disimpan dalam columnPtr.
// Fungsi ini akan terus meminta input ulang jika input tidak valid atau jika input adalah "0" untuk membatalkan.
//
// Parameter:
// - _struct_ (string): String yang menunjukkan jenis struktur yang sedang diproses, bisa "items", "item", "transactions", atau "transaction".
// - prmpt (string): Teks prompt yang menunjukkan apa yang harus dimasukkan pengguna.
// - columnPtr (*string): Pointer ke string tempat nama kolom yang valid akan disimpan.
func InputColumnName(_struct_ string, prmpt string, columnPtr *string) {
	var invalidInputErrText types.Text = types.NewText("Undefined column name. Try again.")
	invalidInputErrText.SetColor("red")

	var stopInput bool = false
	var temp string

	for !stopInput {
		// Menampilkan prompt dan instruksi
		fmt.Print(RightArrowedPrompt(prmpt + " (0 to cancel) "))
		fmt.Scanln(&temp)

		var validInput bool

		// Memeriksa validitas nama kolom berdasarkan struktur
		if _struct_ == "items" || _struct_ == "item" {
			var items types.Items
			validInput = items.IsColumn(temp) || temp == "0"
		} else if _struct_ == "transactions" || _struct_ == "transaction" {
			var transactions types.Transactions
			validInput = transactions.IsColumn(temp) || temp == "0"
		}

		// Jika input valid, keluar dari loop
		if validInput {
			stopInput = true
		} else {
			// Jika input tidak valid, tampilkan pesan error
			fmt.Println(invalidInputErrText.Colored)
		}
	}

	*columnPtr = temp
}
