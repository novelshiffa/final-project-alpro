package types

import (
	"fmt"
	"strings"
)

type Item struct {
	Id       int
	Name     string
	Category string
	Price    int
	Stock    int
}

type Items struct {
	Items  [NMAX]Item
	Length int
}

// IsColumn memeriksa apakah nama kolom yang diberikan valid untuk digunakan dalam struktur Items.
//
// Parameter:
// - columnName (string): Nama kolom yang ingin diperiksa validitasnya.
//
// Return:
// - bool: true jika columnName adalah salah satu dari "id", "name", "category", "price", atau "stock", false jika tidak.
func (items *Items) IsColumn(columnName string) bool {
	lowerCasedColumnName := strings.ToLower(columnName)
	return lowerCasedColumnName == "id" || lowerCasedColumnName == "name" || lowerCasedColumnName == "category" || lowerCasedColumnName == "price" || lowerCasedColumnName == "stock"
}

// getMaxCharOnName mengembalikan panjang maksimum karakter dari nama item dalam struktur Items.
// Fungsi ini digunakan untuk menentukan lebar kolom yang diperlukan saat menampilkan data item.
//
// Return:
// - int: Panjang maksimum karakter dari nama item di dalam struktur Items.
func (items *Items) getMaxCharOnName() int {
	// Algoritma mencari nilai maksimum
	max := 20 // Nilai default untuk nama maksimum yang diharapkan
	for i := 0; i < items.Length; i++ {
		if len(items.Items[i].Name) > max {
			max = len(items.Items[i].Name)
		}
	}
	return max
}

// getMaxId mengembalikan ID maksimum dari semua item dalam struktur Items.
// Fungsi ini digunakan untuk menentukan ID tertinggi di antara semua item yang ada.
//
// Return:
// - int: ID maksimum dari semua item di dalam struktur Items.
func (items *Items) getMaxId() int {
	max := 0

	for i := 0; i < items.Length; i++ {
		if items.Items[i].Id > max {
			max = items.Items[i].Id
		}
	}

	return max
}

// ShowInTable menampilkan semua item dalam struktur Items dalam format tabel yang terformat rapi ke dalam terminal.
// Tabel mencakup kolom ID, Nama, Kategori, Harga, dan Stok untuk setiap item.
// Panjang kolom Nama (Name) disesuaikan secara dinamis berdasarkan panjang maksimum nama item.
func (p *Items) ShowInTable() {
	// Mendapatkan panjang maksimum nama
	nameWidth := p.getMaxCharOnName()

	// Membuat format string header secara dinamis
	headerFormat := fmt.Sprintf("%%-5s %%-%ds %%-15s %%-10s %%-10s\n", nameWidth)
	// Membuat format string untuk setiap baris item
	rowFormat := fmt.Sprintf("%%-5d %%-%ds %%-15s %%-10d %%-10d\n", nameWidth)

	// Header tabel
	fmt.Printf(headerFormat, "ID", "Name", "Category", "Price", "Stock")
	fmt.Println("--------------------------------------------------------------")

	// Mencetak setiap item dalam tabel
	for i := 0; i < p.Length; i++ {
		item := p.Items[i]
		fmt.Printf(rowFormat, item.Id, item.Name, item.Category, item.Price, item.Stock)
	}
}

// FindById mencari item berdasarkan ID menggunakan algoritma pencarian urut (sequential search).
//
// Parameter:
// - id (int): ID dari item yang ingin dicari.
//
// Return:
//   - int: Indeks pertama dimana item dengan ID yang cocok ditemukan dalam struktur Items.
//     Mengembalikan -1 jika item dengan ID yang cocok tidak ditemukan.
func (p *Items) FindById(id int) int {
	// Algoritma pencarian urut (sequential search)
	for i := 0; i < p.Length; i++ {
		if p.Items[i].Id == id {
			return i
		}
	}

	return -1
}

// AddNew menambahkan item baru ke dalam struktur Items.
//
// Parameter:
// - item (Item): Item yang ingin ditambahkan ke dalam struktur Items.
func (p *Items) AddNew(item Item) {
	// Memeriksa apakah panjang Items sudah mencapai batas maksimum
	if p.Length == NMAX {
		panic("Batas maksimum! Tidak bisa menambahkan data.") // Panic jika sudah mencapai batas maksimum
	}

	// Menambahkan item ke dalam array Items dan mengatur ID sesuai dengan urutan penambahan
	p.Items[p.Length] = item
	p.Items[p.Length].Id = p.getMaxId() + 1
	p.Length++
}

// FetchAll mencetak semua item yang ada dalam struktur Items ke terminal.
// Setiap item dicetak dalam baris terpisah.
func (p *Items) FetchAll() {
	for i := 0; i < p.Length; i++ {
		fmt.Println(p.Items[i]) // Mencetak setiap item dalam struktur Items
	}
}

// Delete menghapus item dari struktur Items berdasarkan indeks yang diberikan.
//
// Parameter:
// - idx (int): Indeks dari item yang ingin dihapus dari struktur Items.
func (p *Items) Delete(idx int) {
	// Memeriksa apakah indeks berada dalam rentang yang valid
	if idx < 0 || idx >= p.Length {
		panic("Index out of range.") // Panic jika indeks di luar rentang yang valid
	}

	// Memindahkan semua item setelah indeks yang dihapus ke posisi sebelumnya
	for i := idx; i < p.Length; i++ {
		p.Items[i] = p.Items[i+1]
	}

	p.Length-- // Mengurangi panjang struktur Items setelah penghapusan item
}

// FilterBy melakukan filter terhadap item dalam struktur Items berdasarkan nama kolom dan nilai yang diberikan.
// Fungsi ini mengembalikan struktur Items baru yang berisi item-item yang lolos filter.
//
// Parameter:
// - columnName (string): Nama kolom yang digunakan untuk filter. Harus berupa salah satu dari "id", "name", "category", "price", atau "stock".
// - value (string): Nilai yang digunakan untuk filter. Item-item dalam struktur Items akan dipertahankan jika nilainya sama dengan value sesuai kolom columnName.
//
// Return:
// - Items: Struktur Items baru yang berisi item-item yang lolos filter berdasarkan columnName dan value.
func (p *Items) FilterBy(columnName string, value string) Items {
	columnName = strings.ToLower(columnName)

	var items Items // Struktur Items baru untuk menyimpan item-item yang lolos filter
	columnFilters := map[string](func(item Item) bool){
		"id":       func(item Item) bool { return fmt.Sprintf("%d", item.Id) == value },
		"name":     func(item Item) bool { return item.Name == value },
		"category": func(item Item) bool { return item.Category == value },
		"price":    func(item Item) bool { return fmt.Sprintf("%d", item.Price) == value },
		"stock":    func(item Item) bool { return fmt.Sprintf("%d", item.Stock) == value },
	}

	// Mendapatkan fungsi filter yang sesuai berdasarkan columnName
	filterFunc, exists := columnFilters[columnName]

	// Panic jika nama kolom tidak ditemukan atau tidak valid
	if !exists || !p.IsColumn(columnName) {
		panic("Undefined column.")
	}

	// Melakukan iterasi untuk setiap item dalam struktur Items
	for i := 0; i < p.Length; i++ {
		// Jika item lolos filter, tambahkan ke dalam struktur Items baru
		if filterFunc(p.Items[i]) {
			items.Items[items.Length] = p.Items[i]
			items.Length++
		}
	}

	return items // Mengembalikan struktur Items baru yang telah difilter
}

// SortBy melakukan pengurutan item dalam struktur Items berdasarkan nama kolom dan arah pengurutan yang diberikan.
// Fungsi ini mengembalikan struktur Items baru yang berisi item-item yang sudah terurut.
//
// Parameter:
// - columnName (string): Nama kolom yang digunakan untuk pengurutan. Harus berupa salah satu dari "id", "name", "category", "price", atau "stock".
// - ascending (bool): Jika true, item akan diurutkan secara ascending (terkecil ke terbesar). Jika false, item akan diurutkan secara descending (terbesar ke terkecil).
//
// Return:
// - Items: Struktur Items baru yang berisi item-item yang sudah terurut berdasarkan columnName dan arah pengurutan.
func (p *Items) SortBy(columnName string, ascending bool) Items {
	// Selection Sort

	// Panic jika nama kolom tidak ditemukan atau tidak valid
	if !p.IsColumn(columnName) {
		panic("Undefined column name.")
	}

	columnName = strings.ToLower(columnName)
	var items Items

	// Meng-copy semua item dari p.Items ke items.Items
	for i := 0; i < p.Length; i++ {
		items.Items[i] = p.Items[i]
	}

	items.Length = p.Length

	// Algoritma Selection Sort untuk melakukan pengurutan
	for i := 0; i < items.Length-1; i++ {
		key := i

		// Menemukan indeks item terkecil/terbesar berdasarkan columnName dan ascending
		for j := i + 1; j < p.Length; j++ {
			switch columnName {
			case "id":
				if items.Items[j].Id < items.Items[key].Id && ascending || items.Items[j].Id > items.Items[key].Id && !ascending {
					key = j
				}
			case "name":
				if items.Items[j].Name < items.Items[key].Name && ascending || items.Items[j].Name > items.Items[key].Name && !ascending {
					key = j
				}
			case "category":
				if items.Items[j].Category < items.Items[key].Category && ascending || items.Items[j].Category > items.Items[key].Category && !ascending {
					key = j
				}
			case "price":
				if items.Items[j].Price < items.Items[key].Price && ascending || items.Items[j].Price > items.Items[key].Price && !ascending {
					key = j
				}
			case "stock":
				if items.Items[j].Stock < items.Items[key].Stock && ascending || items.Items[j].Stock > items.Items[key].Stock && !ascending {
					key = j
				}
			}

		}

		// Menukar item pada indeks i dengan item pada indeks key (item terkecil/terbesar)
		items.Items[i], items.Items[key] = items.Items[key], items.Items[i]
	}

	return items // Mengembalikan struktur Items baru yang sudah terurut
}
