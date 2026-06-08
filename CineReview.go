package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Tipe bentukan Film
type Film struct {
	judul     string
	tahun     int
	deskripsi string
	genre     string
	rating    float64
}

// Array statis global
var daftarFilm [100]Film
var jumlahFilm int = 0

// Fungsi bantuan input
func bacaString(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func bacaInt(prompt string) int {
	for {
		fmt.Print(prompt)
		var input int
		_, err := fmt.Scanln(&input)
		if err == nil {
			return input
		}
		fmt.Println("Input tidak valid, masukkan angka.")
	}
}

func bacaFloat(prompt string) float64 {
	for {
		fmt.Print(prompt)
		var input float64
		_, err := fmt.Scanln(&input)
		if err == nil {
			return input
		}
		fmt.Println("Input tidak valid, masukkan angka desimal.")
	}
}

// Prosedur tambah film
func tambahFilm() {
	if jumlahFilm >= 100 {
		fmt.Println("❌ Kapasitas maksimum film tercapai (100).")
		return
	}
	var f Film
	f.judul = bacaString("Judul film: ")
	f.tahun = bacaInt("Tahun rilis: ")
	f.deskripsi = bacaString("Deskripsi singkat: ")
	f.genre = bacaString("Genre: ")
	f.rating = bacaFloat("Rating (1-10): ")
	for f.rating < 1 || f.rating > 10 {
		fmt.Println("Rating harus antara 1 dan 10.")
		f.rating = bacaFloat("Rating (1-10): ")
	}
	daftarFilm[jumlahFilm] = f
	jumlahFilm++
	fmt.Println("✅ Film berhasil ditambahkan!")
}

// Prosedur ubah film (sequential search tanpa break)
func ubahFilm() {
	if jumlahFilm == 0 {
		fmt.Println("Belum ada film.")
		return
	}
	cari := bacaString("Masukkan judul film yang akan diubah: ")
	var eel int = -1
	for i := 0; i < jumlahFilm && eel == -1; i++ {
		if strings.EqualFold(daftarFilm[i].judul, cari) {
			eel = i
		}
	}
	if eel == -1 {
		fmt.Println("❌ Film tidak ditemukan.")
		return
	}
	fmt.Println("Data lama:")
	tampilkanFilm(daftarFilm[eel])
	fmt.Println("Masukkan data baru:")
	var f Film
	f.judul = bacaString("Judul baru: ")
	f.tahun = bacaInt("Tahun rilis baru: ")
	f.deskripsi = bacaString("Deskripsi baru: ")
	f.genre = bacaString("Genre baru: ")
	f.rating = bacaFloat("Rating baru (1-10): ")
	for f.rating < 1 || f.rating > 10 {
		fmt.Println("Rating harus antara 1 dan 10.")
		f.rating = bacaFloat("Rating baru (1-10): ")
	}
	daftarFilm[eel] = f
	fmt.Println("✅ Film berhasil diubah!")
}

// Prosedur hapus film (sequential search tanpa break)
func hapusFilm() {
	if jumlahFilm == 0 {
		fmt.Println("Belum ada film.")
		return
	}
	cari := bacaString("Masukkan judul film yang akan dihapus: ")
	var jebb_24 int = -1
	for i := 0; i < jumlahFilm && jebb_24 == -1; i++ {
		if strings.EqualFold(daftarFilm[i].judul, cari) {
			jebb_24 = i
		}
	}
	if jebb_24 == -1 {
		fmt.Println("❌ Film tidak ditemukan.")
		return
	}
	for i := jebb_24; i < jumlahFilm-1; i++ {
		daftarFilm[i] = daftarFilm[i+1]
	}
	jumlahFilm--
	fmt.Println("✅ Film berhasil dihapus!")
}

// Prosedur tampilkan satu film
func tampilkanFilm(f Film) {
	fmt.Printf("📽️ %s (%d)\n", f.judul, f.tahun)
	fmt.Printf("   Genre: %s | Rating: %.1f\n", f.genre, f.rating)
	fmt.Printf("   Deskripsi: %s\n", f.deskripsi)
}

// Pencarian judul dengan sequential search
func cariJudulSequential() {
	if jumlahFilm == 0 {
		fmt.Println("Belum ada film.")
		return
	}
	cari := bacaString("Masukkan judul yang dicari: ")
	found := false
	for i := 0; i < jumlahFilm; i++ {
		if strings.Contains(strings.ToLower(daftarFilm[i].judul), strings.ToLower(cari)) {
			tampilkanFilm(daftarFilm[i])
			fmt.Println("---")
			found = true
		}
	}
	if !found {
		fmt.Println("❌ Film tidak ditemukan.")
	}
}

// Pencarian judul dengan binary search (tanpa break)
func cariJudulBinary() {
	if jumlahFilm == 0 {
		fmt.Println("Belum ada film.")
		return
	}
	cari := bacaString("Masukkan judul yang dicari (binary search): ")

	// Buat indeks statis lalu urutkan berdasarkan judul
	var indeks [100]int
	for i := 0; i < jumlahFilm; i++ {
		indeks[i] = i
	}
	// Selection sort pada indeks
	for i := 0; i < jumlahFilm-1; i++ {
		minIdx := i
		for j := i + 1; j < jumlahFilm; j++ {
			if daftarFilm[indeks[j]].judul < daftarFilm[indeks[minIdx]].judul {
				minIdx = j
			}
		}
		indeks[i], indeks[minIdx] = indeks[minIdx], indeks[i]
	}

	// Binary search tanpa break
	low, high := 0, jumlahFilm-1
	var eel int = -1
	for low <= high && eel == -1 {
		mid := (low + high) / 2
		cmp := strings.Compare(strings.ToLower(daftarFilm[indeks[mid]].judul), strings.ToLower(cari))
		if cmp == 0 {
			eel = indeks[mid]
		} else if cmp < 0 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if eel == -1 {
		fmt.Println("❌ Film tidak ditemukan.")
	} else {
		tampilkanFilm(daftarFilm[eel])
	}
}

// Pencarian genre dengan sequential search
func cariGenreSequential() {
	if jumlahFilm == 0 {
		fmt.Println("Belum ada film.")
		return
	}
	cari := bacaString("Masukkan genre yang dicari: ")
	found := false
	for i := 0; i < jumlahFilm; i++ {
		if strings.EqualFold(daftarFilm[i].genre, cari) {
			tampilkanFilm(daftarFilm[i])
			fmt.Println("---")
			found = true
		}
	}
	if !found {
		fmt.Println("❌ Tidak ada film dengan genre tersebut.")
	}
}

// Pengurutan dan penampilan (Selection & Insertion Sort)
func urutDanTampilkan() {
	if jumlahFilm == 0 {
		fmt.Println("Belum ada film.")
		return
	}
	fmt.Println("\n=== Urutkan Data Film ===")
	fmt.Println("1. Berdasarkan Rating (Tertinggi ke Terendah) [Selection]")
	fmt.Println("2. Berdasarkan Rating (Terendah ke Tertinggi) [Insertion]")
	fmt.Println("3. Berdasarkan Tahun Rilis (Terbaru ke Terlama) [Selection]")
	fmt.Println("4. Berdasarkan Tahun Rilis (Terlama ke Terbaru) [Insertion]")
	pil := bacaInt("Pilih: ")

	var indeks [100]int
	for i := 0; i < jumlahFilm; i++ {
		indeks[i] = i
	}

	switch pil {
	case 1: // Rating desc (Selection)
		for i := 0; i < jumlahFilm-1; i++ {
			maxIdx := i
			for j := i + 1; j < jumlahFilm; j++ {
				if daftarFilm[indeks[j]].rating > daftarFilm[indeks[maxIdx]].rating {
					maxIdx = j
				}
			}
			indeks[i], indeks[maxIdx] = indeks[maxIdx], indeks[i]
		}
	case 2: // Rating asc (Insertion)
		for i := 1; i < jumlahFilm; i++ {
			key := indeks[i]
			j := i - 1
			for j >= 0 && daftarFilm[indeks[j]].rating > daftarFilm[key].rating {
				indeks[j+1] = indeks[j]
				j--
			}
			indeks[j+1] = key
		}
	case 3: // Tahun desc (Selection)
		for i := 0; i < jumlahFilm-1; i++ {
			maxIdx := i
			for j := i + 1; j < jumlahFilm; j++ {
				if daftarFilm[indeks[j]].tahun > daftarFilm[indeks[maxIdx]].tahun {
					maxIdx = j
				}
			}
			indeks[i], indeks[maxIdx] = indeks[maxIdx], indeks[i]
		}
	case 4: // Tahun asc (Insertion)
		for i := 1; i < jumlahFilm; i++ {
			key := indeks[i]
			j := i - 1
			for j >= 0 && daftarFilm[indeks[j]].tahun > daftarFilm[key].tahun {
				indeks[j+1] = indeks[j]
				j--
			}
			indeks[j+1] = key
		}
	default:
		fmt.Println("Pilihan tidak valid.")
		return
	}

	fmt.Println("\n🎬 Hasil pengurutan:")
	for k := 0; k < jumlahFilm; k++ {
		f := daftarFilm[indeks[k]]
		fmt.Printf("%d. %s (%d) - Rating: %.1f - Genre: %s\n", k+1, f.judul, f.tahun, f.rating, f.genre)
	}
}

// Statistik jumlah film per genre dan rata-rata rating
func tampilkanStatistik() {
	if jumlahFilm == 0 {
		fmt.Println("Belum ada film.")
		return
	}
	type genreCount struct {
		genre string
		count int
	}
	var genres [100]genreCount
	var totalGenre int = 0
	var totalRating float64 = 0.0

	for i := 0; i < jumlahFilm; i++ {
		g := daftarFilm[i].genre
		totalRating += daftarFilm[i].rating
		found := false
		for j := 0; j < totalGenre && !found; j++ {
			if genres[j].genre == g {
				genres[j].count++
				found = true
			}
		}
		if !found {
			genres[totalGenre] = genreCount{g, 1}
			totalGenre++
		}
	}

	avgRating := totalRating / float64(jumlahFilm)
	fmt.Println("\n+++ CineReview +++")
	fmt.Println("📊 STATISTIK KOLEKSI FILM")
	fmt.Println("Jumlah film per genre:")
	for i := 0; i < totalGenre; i++ {
		fmt.Printf("  - %s : %d film\n", genres[i].genre, genres[i].count)
	}
	fmt.Printf("Rata-rata rating seluruh film : %.2f\n", avgRating)
	fmt.Println("+++ CineReview +++")
}

// Menu utama
func main() {
	fmt.Println("🎬 CineReview - Aplikasi Katalog & Rating Film")
	fmt.Println("Kode EEL: CINEREV-001 | Akun @jebb_24")
	for {
		fmt.Println("\n=== MENU UTAMA ===")
		fmt.Println("1. Tambah Film")
		fmt.Println("2. Ubah Film")
		fmt.Println("3. Hapus Film")
		fmt.Println("4. Cari Film (Berdasarkan Judul/Genre)")
		fmt.Println("5. Urutkan & Tampilkan Film")
		fmt.Println("6. Statistik Film")
		fmt.Println("7. Keluar")
		pil := bacaInt("Pilihan: ")

		switch pil {
		case 1:
			tambahFilm()
		case 2:
			ubahFilm()
		case 3:
			hapusFilm()
		case 4:
			fmt.Println("\n--- Pencarian ---")
			fmt.Println("1. Cari Judul (Sequential)")
			fmt.Println("2. Cari Judul (Binary Search)")
			fmt.Println("3. Cari Genre (Sequential)")
			sub := bacaInt("Pilih: ")
			switch sub {
			case 1:
				cariJudulSequential()
			case 2:
				cariJudulBinary()
			case 3:
				cariGenreSequential()
			default:
				fmt.Println("Pilihan tidak valid.")
			}
		case 5:
			urutDanTampilkan()
		case 6:
			tampilkanStatistik()
		case 7:
			fmt.Println("Terima kasih telah menggunakan CineReview!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
