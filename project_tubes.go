package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// ── STRUCT ──────────────────────────────────────────────────
type Warga struct {
	ID     int
	Nama   string
	Alamat string
	HP     string
	Masuk  string
}

type Sampah struct {
	Kode string
	Nama string
}

type Setoran struct {
	ID      int
	WID     int
	Warga   string
	Kode    string
	Jenis   string
	Berat   float64
	Tanggal string
}

// ── DATA GLOBAL ─────────────────────────────────────────────
var (
	warga   []Warga
	setoran []Setoran
	wID     = 1
	sID     = 1
	sc      *bufio.Scanner

	jenisS = []Sampah{
		{"ORG", "Organik (Dapur)"},
		{"PLS", "Plastik"},
		{"KRT", "Kardus / Kertas"},
		{"LOG", "Logam / Kaleng"},
		{"KAC", "Kaca / Botol"},
		{"B3", "B3 (Baterai, dll)"},
	}
)

// ── HELPER ──────────────────────────────────────────────────
func in(prompt string) string {
	fmt.Print(prompt)
	sc.Scan()
	return strings.TrimSpace(sc.Text())
}

func inInt(prompt string) int {
	for {
		v, err := strconv.Atoi(in(prompt))
		if err == nil {
			return v
		}
		fmt.Println("  [!] Masukkan angka.")
	}
}

func inFloat(prompt string) float64 {
	for {
		v, err := strconv.ParseFloat(in(prompt), 64)
		if err == nil && v > 0 {
			return v
		}
		fmt.Println("  [!] Masukkan angka desimal positif. Contoh: 2.5")
	}
}

func cls()            { fmt.Print("\033[H\033[2J") }
func garis()          { fmt.Println("────────────────────────────────────────────────────────────────────────────────────────") }
func jeda()           { in("  Tekan Enter untuk melanjutkan...") }
func ok(s string)     { fmt.Printf("  [✓] %s\n", s) }
func warn(s string)   { fmt.Printf("  [!] %s\n", s) }

func hdr(judul string) {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║       🗑️   WASTE-TRACK  —  Bank Sampah Digital            ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	if judul != "" {
		fmt.Printf("  [ %s ]\n", judul)
		garis()
	}
}

// ── UTIL WARGA ───────────────────────────────────────────────
func cariWarga(id int) int {
	for i, w := range warga {
		if w.ID == id {
			return i
		}
	}
	return -1
}

func tabelWarga(list []Warga) {
	if len(list) == 0 {
		warn("Belum ada data warga.")
		return
	}
	fmt.Printf("  %-4s %-22s %-28s %-14s %-12s\n", "ID", "Nama", "Alamat", "NO HP", "Masuk")
	garis()
	for _, w := range list {
		fmt.Printf("  %-4d %-22s %-28s %-14s %-12s\n", w.ID, w.Nama, w.Alamat, w.HP, w.Masuk)
	}
}

//CRUD WARGA
func tambah() {
	cls(); hdr("TAMBAH WARGA")
	nama := in("  Nama   : ")
	if nama == "" {
		warn("Nama tidak boleh kosong.")
		return
	}
	alamat := in("  Alamat : ")
	hp := in("  HP     : ")
	w := Warga{wID, nama, alamat, hp, time.Now().Format("02-01-2006")}
	warga = append(warga, w)
	wID++
	garis()
	ok(fmt.Sprintf("Warga '%s' ditambahkan (ID: %d)", nama, w.ID))
}

func lihatSemua() {
	cls(); hdr("DAFTAR WARGA")
	fmt.Printf("  Total: %d warga\n\n", len(warga))
	tabelWarga(warga)
}

func detail() {
	cls(); hdr("DETAIL WARGA")
	if len(warga) == 0 {
		warn("Belum ada data warga.")
		return
	}
	tabelWarga(warga)
	garis()
	id := inInt("  ID warga: ")
	idx := cariWarga(id)
	if idx == -1 {
		warn(fmt.Sprintf("ID %d tidak ditemukan.", id))
		return
	}
	w := warga[idx]
	garis()
	fmt.Printf("  ID     : %d\n", w.ID)
	fmt.Printf("  Nama   : %s\n", w.Nama)
	fmt.Printf("  Alamat : %s\n", w.Alamat)
	fmt.Printf("  HP     : %s\n", w.HP)
	fmt.Printf("  Masuk  : %s\n", w.Masuk)
}

func ubah() {
	cls(); hdr("UBAH WARGA")
	if len(warga) == 0 {
		warn("Belum ada data warga.")
		return
	}
	tabelWarga(warga)
	garis()
	id := inInt("  ID warga yang diubah: ")
	idx := cariWarga(id)
	if idx == -1 {
		warn(fmt.Sprintf("ID %d tidak ditemukan.", id))
		return
	}
	w := warga[idx]
	fmt.Printf("\n  Data: %s | %s | %s\n", w.Nama, w.Alamat, w.HP)
	garis()
	fmt.Println("  (Kosongkan = tidak diubah)")
	if v := in("  Nama baru   : "); v != "" {
		warga[idx].Nama = v
	}
	if v := in("  Alamat baru : "); v != "" {
		warga[idx].Alamat = v
	}
	if v := in("  NO HP baru  : "); v != "" {
		warga[idx].HP = v
	}
	for i, s := range setoran {
		if s.WID == id {
			setoran[i].Warga = warga[idx].Nama
		}
	}
	garis(); ok("Data warga diperbarui.")
}

func hapus() {
	cls(); hdr("HAPUS WARGA")
	if len(warga) == 0 {
		warn("Belum ada data warga.")
		return
	}
	tabelWarga(warga)
	garis()
	id := inInt("  ID warga yang dihapus: ")
	idx := cariWarga(id)
	if idx == -1 {
		warn(fmt.Sprintf("ID %d tidak ditemukan.", id))
		return
	}
	fmt.Printf("  Nama: %s\n", warga[idx].Nama)
	if strings.ToLower(in("  Yakin hapus? (y/n): ")) == "y" {
		nama := warga[idx].Nama
		warga = append(warga[:idx], warga[idx+1:]...)
		var sisa []Setoran
		for _, s := range setoran {
			if s.WID != id {
				sisa = append(sisa, s)
			}
		}
		setoran = sisa
		ok(fmt.Sprintf("Warga '%s' dihapus.", nama))
	} else {
		fmt.Println("  [x] Dibatalkan.")
	}
}

func menuWarga() {
	for {
		cls(); hdr("MANAJEMEN WARGA")
		fmt.Printf("  Total: %d warga\n\n", len(warga))
		fmt.Println("  1. Tambah Warga")
		fmt.Println("  2. Lihat Semua")
		fmt.Println("  3. Detail Warga")
		fmt.Println("  4. Ubah Warga")
		fmt.Println("  5. Hapus Warga")
		fmt.Println("  0. Kembali")
		garis()
		switch in("  Pilihan: ") {
		case "1":
			tambah()
		case "2":
			lihatSemua()
		case "3":
			detail()
		case "4":
			ubah()
		case "5":
			hapus()
		case "0":
			return
		default:
			warn("Pilihan tidak valid.")
		}
		fmt.Println(); jeda()
	}
}

// ── B. SETORAN SAMPAH ────────────────────────────────────────
func tabelSetoran(list []Setoran) {
	if len(list) == 0 {
		warn("Tidak ada setoran.")
		return
	}
	fmt.Printf("  %-4s %-20s %-20s %-8s %-12s\n", "ID", "Warga", "Jenis", "Berat", "Tanggal")
	garis()
	for _, s := range list {
		fmt.Printf("  %-4d %-20s %-20s %-8.2f %-12s\n", s.ID, s.Warga, s.Jenis, s.Berat, s.Tanggal)
	}
}

func catatSetoran() {
	cls(); hdr("CATAT SETORAN")
	if len(warga) == 0 {
		warn("Tambah warga dulu.")
		return
	}
	tabelWarga(warga)
	garis()
	wid := inInt("  Pilih ID Warga: ")
	idx := cariWarga(wid)
	if idx == -1 {
		warn(fmt.Sprintf("ID %d tidak ditemukan.", wid))
		return
	}
	fmt.Printf("  Warga: %s\n\n", warga[idx].Nama)
	for i, j := range jenisS {
		fmt.Printf("  %d. [%s] %s\n", i+1, j.Kode, j.Nama)
	}
	garis()
	pil := inInt("  Pilih jenis (nomor): ")
	if pil < 1 || pil > len(jenisS) {
		warn("Pilihan tidak valid.")
		return
	}
	j := jenisS[pil-1]
	berat := inFloat("  Berat (kg): ")
	tgl := in("  Tanggal (DD-MM-YYYY, kosong=hari ini): ")
	if tgl == "" {
		tgl = time.Now().Format("02-01-2006")
	}
	s := Setoran{sID, wid, warga[idx].Nama, j.Kode, j.Nama, berat, tgl}
	setoran = append(setoran, s)
	sID++
	garis(); ok(fmt.Sprintf("%.2f kg %s dari '%s' dicatat.", berat, j.Nama, warga[idx].Nama))
}

func hapusSetoran() {
	cls(); hdr("HAPUS SETORAN")
	if len(setoran) == 0 {
		warn("Belum ada setoran.")
		return
	}
	tabelSetoran(setoran)
	garis()
	id := inInt("  ID setoran yang dihapus: ")
	for i, s := range setoran {
		if s.ID == id {
			fmt.Printf("  %s | %s | %.2f kg\n", s.Warga, s.Jenis, s.Berat)
			if strings.ToLower(in("  Yakin? (y/n): ")) == "y" {
				setoran = append(setoran[:i], setoran[i+1:]...)
				ok("Setoran dihapus.")
			} else {
				fmt.Println("  [x] Dibatalkan.")
			}
			return
		}
	}
	warn(fmt.Sprintf("ID %d tidak ditemukan.", id))
}

func menuSetoran() {
	for {
		cls(); hdr("SETORAN SAMPAH")
		fmt.Printf("  Total setoran: %d\n\n", len(setoran))
		fmt.Println("  1. Catat Setoran Baru")
		fmt.Println("  2. Lihat Semua Setoran")
		fmt.Println("  3. Hapus Setoran")
		fmt.Println("  0. Kembali")
		garis()
		switch in("  Pilihan: ") {
		case "1":
			catatSetoran()
		case "2":
			cls(); hdr("SEMUA SETORAN")
			tabelSetoran(setoran)
		case "3":
			hapusSetoran()
		case "0":
			return
		default:
			warn("Pilihan tidak valid.")
		}
		fmt.Println(); jeda()
	}
}

// ── C. PENCARIAN ─────────────────────────────────────────────
func totalBerat(wid int) float64 {
	total := 0.0
	for _, s := range setoran {
		if s.WID == wid {
			total += s.Berat
		}
	}
	return total
}

func seqNama(kata string) []Warga {
	kata = strings.ToLower(kata)
	var hasil []Warga
	for i, w := range warga {
		fmt.Printf("  → Langkah %d: cek '%s'\n", i+1, w.Nama)
		if strings.Contains(strings.ToLower(w.Nama), kata) {
			hasil = append(hasil, w)
		}
	}
	return hasil
}

func seqID(id int) *Warga {
	for i, w := range warga {
		fmt.Printf("  → Langkah %d: cek ID %d\n", i+1, w.ID)
		if w.ID == id {
			return &warga[i]
		}
	}
	return nil
}

func binID(id int) *Warga {
	// Salin & urutkan by ID
	sorted := make([]Warga, len(warga))
	copy(sorted, warga)
	// insertion sort by ID
	for i := 1; i < len(sorted); i++ {
		key := sorted[i]
		j := i - 1
		for j >= 0 && sorted[j].ID > key.ID {
			sorted[j+1] = sorted[j]
			j--
		}
		sorted[j+1] = key
	}
	lo, hi, langkah := 0, len(sorted)-1, 0
	for lo <= hi {
		langkah++
		mid := (lo + hi) / 2
		fmt.Printf("  → Langkah %d: cek indeks %d (ID=%d)\n", langkah, mid, sorted[mid].ID)
		if sorted[mid].ID == id {
			for i := range warga {
				if warga[i].ID == id {
					return &warga[i]
				}
			}
		} else if sorted[mid].ID < id {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return nil
}

func menuCari() {
	for {
		cls(); hdr("PENCARIAN WARGA")
		fmt.Println("  1. Cari Nama   — Sequential Search")
		fmt.Println("  2. Cari ID     — Sequential Search")
		fmt.Println("  3. Cari ID     — Binary Search")
		fmt.Println("  0. Kembali")
		garis()
		switch in("  Pilihan: ") {
		case "1":
			cls(); hdr("SEQUENTIAL — NAMA")
			kata := in("  Kata kunci: ")
			garis()
			hasil := seqNama(kata)
			garis()
			if len(hasil) == 0 {
				warn("Tidak ditemukan.")
			} else {
				fmt.Printf("  Ditemukan %d warga:\n\n", len(hasil))
				tabelWarga(hasil)
			}
		case "2":
			cls(); hdr("SEQUENTIAL — ID")
			id := inInt("  ID warga: ")
			garis()
			w := seqID(id)
			garis()
			if w == nil {
				warn(fmt.Sprintf("ID %d tidak ditemukan.", id))
			} else {
				fmt.Printf("  ID: %d | Nama: %s | HP: %s\n", w.ID, w.Nama, w.HP)
			}
		case "3":
			cls(); hdr("BINARY SEARCH — ID")
			if len(warga) == 0 {
				warn("Belum ada data warga.")
				break
			}
			id := inInt("  ID warga: ")
			garis()
			w := binID(id)
			garis()
			if w == nil {
				warn(fmt.Sprintf("ID %d tidak ditemukan.", id))
			} else {
				fmt.Printf("  ID: %d | Nama: %s | HP: %s\n", w.ID, w.Nama, w.HP)
			}
		case "0":
			return
		default:
			warn("Pilihan tidak valid.")
		}
		fmt.Println(); jeda()
	}
}

// ── D. PENGURUTAN ────────────────────────────────────────────
type WB struct {
	Warga
	Berat float64
}

func buatWB() []WB {
	var list []WB
	for _, w := range warga {
		list = append(list, WB{w, totalBerat(w.ID)})
	}
	return list
}

func tabelWB(list []WB) {
	fmt.Printf("  %-4s %-22s %-10s\n", "ID", "Nama", "Total(kg)")
	garis()
	for i, wb := range list {
		fmt.Printf("  %-4d %-22s %-10.2f  (peringkat %d)\n", wb.ID, wb.Nama, wb.Berat, i+1)
	}
}

func selSort(list []WB) []WB {
	n := len(list)
	for i := 0; i < n-1; i++ {
		max := i
		for j := i + 1; j < n; j++ {
			if list[j].Berat > list[max].Berat {
				max = j
			}
		}
		list[i], list[max] = list[max], list[i]
		fmt.Printf("  → Pass %d: '%s' ke posisi %d\n", i+1, list[i].Nama, i+1)
	}
	return list
}

func insSort(list []WB) []WB {
	for i := 1; i < len(list); i++ {
		key := list[i]
		j := i - 1
		for j >= 0 && list[j].Berat < key.Berat {
			list[j+1] = list[j]
			j--
		}
		list[j+1] = key
		fmt.Printf("  → Sisip '%s' ke posisi %d\n", key.Nama, j+2)
	}
	return list
}

func menuUrut() {
	for {
		cls(); hdr("PENGURUTAN DATA")
		fmt.Println("  Urut berdasarkan total berat setoran (terbanyak dulu)")
		garis()
		fmt.Println("  1. Selection Sort")
		fmt.Println("  2. Insertion Sort")
		fmt.Println("  0. Kembali")
		garis()
		switch in("  Pilihan: ") {
		case "1":
			cls(); hdr("SELECTION SORT")
			list := buatWB()
			garis()
			list = selSort(list)
			garis()
			tabelWB(list)
		case "2":
			cls(); hdr("INSERTION SORT")
			list := buatWB()
			garis()
			list = insSort(list)
			garis()
			tabelWB(list)
		case "0":
			return
		default:
			warn("Pilihan tidak valid.")
		}
		fmt.Println(); jeda()
	}
}

// ── E. STATISTIK MINGGUAN ────────────────────────────────────
func mingguDari(tgl string) (time.Time, time.Time, error) {
	t, err := time.Parse("02-01-2006", tgl)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	wd := int(t.Weekday())
	if wd == 0 {
		wd = 7
	}
	senin := t.AddDate(0, 0, -(wd - 1))
	return senin, senin.AddDate(0, 0, 6), nil
}

func dalamMinggu(tgl string, senin, minggu time.Time) bool {
	t, err := time.Parse("02-01-2006", tgl)
	if err != nil {
		return false
	}
	return !t.Before(senin) && !t.After(minggu)
}

func statistik() {
	cls(); hdr("STATISTIK MINGGUAN")
	ref := in("  Tanggal referensi (DD-MM-YYYY, kosong=hari ini): ")
	if ref == "" {
		ref = time.Now().Format("02-01-2006")
	}
	senin, minggu, err := mingguDari(ref)
	if err != nil {
		warn("Format tanggal tidak valid.")
		return
	}
	fmt.Printf("\n  Periode: %s — %s\n\n", senin.Format("02-01-2006"), minggu.Format("02-01-2006"))

	var mingguIni []Setoran
	for _, s := range setoran {
		if dalamMinggu(s.Tanggal, senin, minggu) {
			mingguIni = append(mingguIni, s)
		}
	}
	if len(mingguIni) == 0 {
		warn("Tidak ada setoran minggu ini.")
		return
	}

	total := 0.0
	perJenis := map[string]float64{}
	perWarga := map[int]float64{}
	namaMap := map[int]string{}
	for _, s := range mingguIni {
		total += s.Berat
		perJenis[s.Jenis] += s.Berat
		perWarga[s.WID] += s.Berat
		namaMap[s.WID] = s.Warga
	}

	garis()
	fmt.Printf("  Total transaksi   : %d\n", len(mingguIni))
	fmt.Printf("  Warga aktif       : %d\n", len(perWarga))
	fmt.Printf("  Total berat       : %.2f kg\n", total)
	fmt.Printf("  Rata-rata/warga   : %.2f kg\n", total/float64(len(perWarga)))
	garis()

	fmt.Println("\n  PER JENIS SAMPAH:")
	fmt.Printf("  %-22s %-10s %s\n", "Jenis", "Berat(kg)", "%")
	garis()
	for jenis, berat := range perJenis {
		pct := berat / total * 100
		bar := strings.Repeat("█", int(pct/5))
		fmt.Printf("  %-22s %-10.2f %5.1f%% %s\n", jenis, berat, pct, bar)
	}

	fmt.Println("\n  PER WARGA:")
	fmt.Printf("  %-4s %-22s %-10s\n", "ID", "Nama", "Berat(kg)")
	garis()
	best, bestBerat := 0, 0.0
	for wid, berat := range perWarga {
		fmt.Printf("  %-4d %-22s %.2f kg\n", wid, namaMap[wid], berat)
		if berat > bestBerat {
			bestBerat = berat
			best = wid
		}
	}
	garis()
	fmt.Printf("  🏆 Terbanyak: %s (%.2f kg)\n", namaMap[best], bestBerat)
}

func menuStatistik() {
	for {
		cls(); hdr("STATISTIK MINGGUAN")
		fmt.Println("  1. Lihat Statistik Minggu Ini")
		fmt.Println("  2. Lihat Semua Setoran")
		fmt.Println("  0. Kembali")
		garis()
		switch in("  Pilihan: ") {
		case "1":
			statistik()
		case "2":
			cls(); hdr("SEMUA SETORAN")
			tabelSetoran(setoran)
		case "0":
			return
		default:
			warn("Pilihan tidak valid.")
		}
		fmt.Println(); jeda()
	}
}

// ── MENU UTAMA ───────────────────────────────────────────────
func menu() {
	for {
		cls(); hdr("")
		fmt.Println("  [ MENU UTAMA ]")
		fmt.Println()
		fmt.Println("  A. Manajemen Warga      (CRUD)")
		fmt.Println("  B. Setoran Sampah       (Catat & Lihat)")
		fmt.Println("  C. Pencarian Warga      (Sequential & Binary)")
		fmt.Println("  D. Pengurutan Data      (Selection & Insertion)")
		fmt.Println("  E. Statistik Mingguan")
		garis()
		fmt.Println("  0. Keluar")
		garis()
		switch strings.ToUpper(in("  Pilihan: ")) {
		case "A":
			menuWarga()
		case "B":
			menuSetoran()
		case "C":
			menuCari()
		case "D":
			menuUrut()
		case "E":
			menuStatistik(); fmt.Println(); jeda()
		case "0":
			cls()
			fmt.Println("  Terima kasih! Jaga kebersihan lingkungan. 🌿")
			fmt.Println()
			os.Exit(0)
		default:
			warn("Pilihan tidak valid.")
			jeda()
		}
	}
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	menu()
}