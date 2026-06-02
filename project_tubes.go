package main

import "fmt"

func manajemen() {
	var y string
for {
		fmt.Println("  [ MANAJEMEN DATA WARGA ]")
		fmt.Printf("  Total warga terdaftar:\n")
		fmt.Println("  1. Tambah Warga Baru")
		fmt.Println("  2. Lihat Semua Warga")
		fmt.Println("  3. Lihat Detail Warga")
		fmt.Println("  4. Ubah Data Warga")
		fmt.Println("  5. Hapus Data Warga")
		fmt.Println("  0. Kembali ke Menu Utama")
		fmt.Print("Pilih (A,B,C,D,E,atau 0): ")
		fmt.Scan(&y)


		switch y {
		case "1":
			tambahWarga()
		case "2":
			tampilkanSemuaWarga()
		case "3":
			detailWarga()
		case "4":
			ubahWarga()
		case "5":
			hapusWarga()
		case "0":
			return
		default:
			fmt.Println("  [!] Pilihan tidak valid.")
		}
	}
}
func setoran() {

}
func Pencarian() {

}
func Pengurutan() {

}

func menu() {
	var x string
	for {
		fmt.Println("  [ MENU UTAMA ]")
		fmt.Println(" ")
		fmt.Println("  A. Manajemen Data Warga  (CRUD)")
		fmt.Println("  B. Pencatatan Setoran Sampah")
		fmt.Println("  C. Pencarian Data Warga")
		fmt.Println("  D. Pengurutan Data")
		fmt.Println("  E. Statistik Mingguan")
		fmt.Println("  0. Keluar")
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Printf("Pilih (A,B,C,D,E,atau 0): ")
		fmt.Scan(&x)
		switch x {
		case "A", "a":
			manajemen()
		case "B", "b":
			setoran()
		case "C", "c":
			Pencarian()
		case "D", "d":
			Pengurutan()
		default:
			fmt.Println("  [!] Pilihan tidak valid.")
		}
	}

}
func main() {
	menu()
}
