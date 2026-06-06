package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	scanner *bufio.Scanner
)

func input(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func menu() {
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
		pilihan := input("Pilih (A,B, atau 0): ")
		pilihan = strings.ToUpper(pilihan)
		switch pilihan {
		case "A":
			manajemen()
		case "B":
			setoran()
		case "C":
			Pencarian()
		case "D":
			Pengurutan()
		case "0":
			break
		default:
			fmt.Println("  [!] Pilihan tidak valid.")
		}
	}

}
func manajemen() {
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
		pilihan := input("Pilih: ")

		switch pilihan {
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
func tambahWarga(){}
func tampilkanSemuaWarga(){}
func detailWarga(){}
func ubahWarga(){}
func hapusWarga(){}



func setoran() {}
func Pencarian() {}
func Pengurutan() {}

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	menu()
}
