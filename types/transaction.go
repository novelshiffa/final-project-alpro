package types

import (
	"fmt"
	"strings"
	"time"
)

type Transaction struct {
	Id       int
	Time     time.Time
	Type     string
	Item     Item
	Quantity int
}

type Transactions struct {
	Items  [NMAX]Transaction
	Length int
}

// IsColumn memeriksa apakah nama kolom yang diberikan adalah kolom yang valid dalam struct Transactions.
//
// Paramater:
// - columnName (string): nama kolom yang akan diperiksa.
//
// Return:
// - bool: true jika columnName adalah "id", "time", "type", "itemid", atau "quantity", false jika tidak.
func (t *Transactions) IsColumn(columnName string) bool {
	lowerCasedColumnName := strings.ToLower(columnName)
	return lowerCasedColumnName == "id" || lowerCasedColumnName == "time" || lowerCasedColumnName == "type" || lowerCasedColumnName == "itemid" || lowerCasedColumnName == "quantity"
}

// isIdSorted memeriksa apakah daftar transaksi diurutkan berdasarkan Id.
//
// Parameter:
//   - ascendingly (bool): boolean yang menentukan urutan pengurutan. Jika true, transaksi harus diurutkan secara menaik (ascending),
//     jika false, transaksi harus diurutkan secara menurun (descending).
//
// Return:
// - bool: true jika daftar transaksi diurutkan sesuai dengan parameter 'ascendingly', false jika tidak.
func (transactions *Transactions) isIdSorted(ascendingly bool) bool {
	for i := 1; i < transactions.Length; i++ {
		if (ascendingly && transactions.Items[i].Id < transactions.Items[i-1].Id) ||
			(!ascendingly && transactions.Items[i].Id > transactions.Items[i-1].Id) {
			return false
		}
	}
	return true
}

// ShowInTable menampilkan daftar transaksi dalam format tabel.
//
// Parameter:
// - Tidak ada parameter.
//
// Return:
// - Tidak ada nilai yang dikembalikan.
func (transactions *Transactions) ShowInTable() {
	fmt.Printf("%-5s %-20s %-10s %-20s %-10s %-10s\n", "ID", "Time", "Type", "Item Id", "Quantity", "Price")
	fmt.Println("-------------------------------------------------------------------------------------------")
	for i := 0; i < transactions.Length; i++ {
		t := transactions.Items[i]
		fmt.Printf("%-5d %-20s %-10s %-20d %-10d %-10.2d\n", t.Id, t.Time.Format("2006-01-02 15:04:05"), t.Type, t.Item.Id, t.Quantity, t.Item.Price)
	}
}

// FindById melakukan pencarian biner (binary search) untuk mencari indeks transaksi dengan Id yang cocok.
// Jika array tidak terurut berdasarkan Id (ascendingly), fungsi akan mengurutkannya secara otomatis sebelum melakukan pencarian.
//
// Parameter:
// - id (int): Id transaksi yang ingin dicari.
//
// Return:
// - int: Indeks dari transaksi dengan Id yang cocok, atau -1 jika tidak ditemukan.
func (t *Transactions) FindById(id int) int {
	// Algoritma pencarian biner (binary search), secara otomatis mengurutkan ascendingly berdasarkan id jika belum terurut

	if !t.isIdSorted(true) {
		t.SortBy("id", true)
	}

	low, high := 0, t.Length-1
	for low <= high {
		mid := (low + high) / 2
		if t.Items[mid].Id == id {
			return mid
		} else if t.Items[mid].Id < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// CreateNew menambahkan transaksi baru ke dalam array Transactions.
// Prosedur ini akan menyebabkan panic jika panjang maksimum (NMAX) sudah tercapai.
//
// Parameter:
// - transaction (Transaction): transaksi yang ingin ditambahkan ke dalam array.
func (t *Transactions) CreateNew(transaction Transaction) {
	if t.Length == NMAX {
		panic("Max length reached.")
	}

	t.Items[t.Length] = transaction
	t.Items[t.Length].Id = t.Length + 1
	t.Length++
}

// FetchAll mencetak semua transaksi dalam array Transactions ke dalam output.
// Setiap transaksi dicetak dalam format yang dimiliki oleh struct Transaction.
//
// Parameter:
// - Tidak ada parameter.
//
// Return:
// - Tidak ada nilai yang dikembalikan.
func (t *Transactions) FetchAll() {
	for i := 0; i < t.Length; i++ {
		fmt.Println(t.Items[i])
	}
}

// Delete menghapus transaksi pada indeks tertentu dari array Transactions.
// Jika indeks yang diberikan berada di luar rentang yang valid, fungsi akan menyebabkan panic.
//
// Parameter:
// - idx (int): Indeks transaksi yang ingin dihapus dari array.
func (t *Transactions) Delete(idx int) {
	// Memeriksa apakah indeks di luar rentang yang valid.
	if idx < 0 || idx >= t.Length {
		panic("Index out of range.")
	}

	// Menggeser semua elemen setelah indeks yang akan dihapus ke kiri.
	for i := idx; i < t.Length-1; i++ {
		t.Items[i] = t.Items[i+1]
	}

	// Mengurangi panjang array Transactions setelah penghapusan.
	t.Length--
}

// FilterBy melakukan filter terhadap array Transactions berdasarkan nama kolom dan nilai yang diberikan.
//
// Parameter:
// - columnName (string): Nama kolom yang digunakan sebagai kriteria filter.
// - value (string): Nilai yang digunakan sebagai kriteria untuk melakukan filter.
//
// Return:
// - Transactions: Transaksi baru yang berisi transaksi yang sudah difilter berdasarkan kolom dan nilai yang diberikan.
func (t *Transactions) FilterBy(columnName string, value string) Transactions {
	// Konversi `columnName` ke huruf kecil.
	columnName = strings.ToLower(columnName)

	// Membuat variabel `transactions` untuk menyimpan hasil filter.
	var transactions Transactions

	// Menyiapkan map `columnFilters` yang berisi fungsi filter untuk setiap kolom yang mungkin.
	columnFilters := map[string](func(transaction Transaction) bool){
		"id": func(transaction Transaction) bool { return fmt.Sprintf("%d", transaction.Id) == value },
		"time": func(transaction Transaction) bool {
			return transaction.Time.Format("2006-01-02 15:04:05") == value || strings.Split(transaction.Time.Format("2006-01-02 15:04:05"), " ")[0] == value
		},
		"type":     func(transaction Transaction) bool { return transaction.Type == value },
		"itemid":   func(transaction Transaction) bool { return fmt.Sprintf("%d", transaction.Item.Id) == value },
		"quantity": func(transaction Transaction) bool { return fmt.Sprintf("%d", transaction.Quantity) == value },
	}

	// Memeriksa apakah `columnName` ada di dalam `columnFilters` dan apakah valid dengan memanggil `IsColumn`.
	filterFunc, exists := columnFilters[columnName]
	if !exists || !t.IsColumn(columnName) {
		panic("Undefined column.")
	}

	// Iterasi melalui setiap transaksi dalam array `t.Items`.
	for i := 0; i < t.Length; i++ {
		// Jika transaksi memenuhi kriteria filter (`filterFunc`), maka tambahkan transaksi tersebut ke dalam `transactions`.
		if filterFunc(t.Items[i]) {
			transactions.Items[transactions.Length] = t.Items[i]
			transactions.Length++
		}
	}

	// Mengembalikan `transactions` yang sudah difilter.
	return transactions
}

// SortBy mengurutkan transaksi dalam array Transactions berdasarkan kolom yang ditentukan secara ascending atau descending menggunakan algoritma insertion sort.
// Fungsi ini akan menyebabkan panic jika nama kolom yang diberikan tidak terdefinisi.
//
// Parameter:
// - columnName (string): Nama kolom yang digunakan sebagai kunci untuk pengurutan (id, time, itemid, type, quantity).
// - ascending (bool): Menentukan apakah pengurutan dilakukan secara ascending (true) atau descending (false).
//
// Return:
// - Transactions: Transaksi baru yang sudah diurutkan berdasarkan kolom dan urutan yang ditentukan.
func (t *Transactions) SortBy(columnName string, ascending bool) Transactions {
	// Memeriksa apakah nama kolom yang diberikan valid.
	if !t.IsColumn(columnName) {
		panic("Undefined column name.")
	}

	// Konversi nama kolom ke huruf kecil untuk konsistensi.
	columnName = strings.ToLower(columnName)

	// Menyiapkan variabel transactions untuk menyimpan hasil pengurutan.
	var transactions Transactions

	// Mengkopi semua elemen dari t ke transactions.
	for i := 0; i < t.Length; i++ {
		transactions.Items[i] = t.Items[i]
	}
	transactions.Length = t.Length

	// Algoritma insertion sort
	for i := 1; i < t.Length; i++ {
		j := i

		// Membandingkan dan menukar elemen sampai posisi yang tepat dalam urutan.
		for j > 0 {
			switch columnName {
			case "id":
				if (transactions.Items[j].Id < transactions.Items[j-1].Id && ascending) ||
					(transactions.Items[j].Id > transactions.Items[j-1].Id && !ascending) {
					transactions.Items[j-1], transactions.Items[j] = transactions.Items[j], transactions.Items[j-1]
				}
			case "time":
				if (transactions.Items[j].Time.Before(transactions.Items[j-1].Time) && ascending) ||
					(transactions.Items[j].Time.After(transactions.Items[j-1].Time) && !ascending) {
					transactions.Items[j-1], transactions.Items[j] = transactions.Items[j], transactions.Items[j-1]
				}
			case "itemid":
				if (transactions.Items[j].Item.Id < transactions.Items[j-1].Item.Id && ascending) ||
					(transactions.Items[j].Item.Id > transactions.Items[j-1].Item.Id && !ascending) {
					transactions.Items[j-1], transactions.Items[j] = transactions.Items[j], transactions.Items[j-1]
				}
			case "type":
				if (transactions.Items[j].Type < transactions.Items[j-1].Type && ascending) ||
					(transactions.Items[j].Type > transactions.Items[j-1].Type && !ascending) {
					transactions.Items[j-1], transactions.Items[j] = transactions.Items[j], transactions.Items[j-1]
				}
			case "quantity":
				if (transactions.Items[j].Quantity < transactions.Items[j-1].Quantity && ascending) ||
					(transactions.Items[j].Quantity > transactions.Items[j-1].Quantity && !ascending) {
					transactions.Items[j-1], transactions.Items[j] = transactions.Items[j], transactions.Items[j-1]
				}
			}

			j--
		}
	}

	// Mengembalikan transactions yang sudah diurutkan.
	return transactions
}
