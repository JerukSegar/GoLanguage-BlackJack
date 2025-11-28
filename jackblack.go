package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const maksKartu = 10

var kartu = [13]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
var nilaiKartu = map[string]int{"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10, "J": 10, "Q": 10, "K": 10, "A": 11}

type Statistik struct {
	Menang int
	Kalah  int
	Seri   int
}

func ambilKartu() string {
	return kartu[rand.Intn(len(kartu))]
}

func hitungNilaiTangan(tangan [maksKartu]string, jumlah int) int {
	var total int
	var jumlahAs int

	total = 0
	jumlahAs = 0

	for i := 0; i < jumlah; i++ {
		var k string
		k = tangan[i]
		total = total + nilaiKartu[k]
		if k == "A" {
			jumlahAs = jumlahAs + 1
		}
	}

	for total > 21 && jumlahAs > 0 {
		total = total - 10
		jumlahAs = jumlahAs - 1
	}

	return total
}

func tampilkanKartu(pemain []string, bandar []string, tunjukkanSemuaBandar bool) {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("                  PERMAINAN BLACKJACK")
	fmt.Println(strings.Repeat("=", 50))

	fmt.Print("Bandar: ")
	if tunjukkanSemuaBandar {
		var tanganBandarArr [maksKartu]string
		copy(tanganBandarArr[:], bandar)
		fmt.Println(bandar, "Total:", hitungNilaiTangan(tanganBandarArr, len(bandar)))
	} else {
		fmt.Printf("[%s, ?]\n", bandar[0])
	}

	var tanganPemainArr [maksKartu]string
	copy(tanganPemainArr[:], pemain)
	fmt.Printf("Anda:   %v Total: %d\n", pemain, hitungNilaiTangan(tanganPemainArr, len(pemain)))
	fmt.Println(strings.Repeat("-", 50))
}

func tampilkanMenu(stat Statistik) {
	fmt.Println("\n" + strings.Repeat("=", 40))
	fmt.Println("           MENU BLACKJACK")
	fmt.Println(strings.Repeat("=", 40))
	fmt.Printf("Statistik: %d Menang | %d Kalah | %d Seri\n", stat.Menang, stat.Kalah, stat.Seri)
	fmt.Println(strings.Repeat("-", 40))
	fmt.Println("1. Mulai Permainan Baru")
	fmt.Println("2. Keluar")
	fmt.Print("Pilih menu (1-2): ")
}

func mulaiPermainan() ([maksKartu]string, int, [maksKartu]string, int) {
	var tanganPemain [maksKartu]string
	var tanganBandar [maksKartu]string

	tanganPemain[0] = ambilKartu()
	tanganPemain[1] = ambilKartu()
	tanganBandar[0] = ambilKartu()
	tanganBandar[1] = ambilKartu()

	return tanganPemain, 2, tanganBandar, 2
}

func giliranPemain(tangan *[maksKartu]string, jumlah *int) {
	for {
		fmt.Print("Hit / Stand: ")
		var pilihan string
		fmt.Scanln(&pilihan)
		pilihan = strings.ToLower(strings.TrimSpace(pilihan))

		if pilihan == "hit" {
			if *jumlah >= maksKartu {
				fmt.Println("Anda sudah memiliki terlalu banyak kartu!")
				break
			}
			tangan[*jumlah] = ambilKartu()
			*jumlah = *jumlah + 1

			var sliceTangan []string
			sliceTangan = tangan[:*jumlah]
			var tanganArr [maksKartu]string
			copy(tanganArr[:], sliceTangan)

			fmt.Println("Kartu Anda:", sliceTangan, "Total:", hitungNilaiTangan(tanganArr, *jumlah))
			if hitungNilaiTangan(tanganArr, *jumlah) > 21 {
				fmt.Println("Bust!")
				return
			}
		} else if pilihan == "stand" {
			break
		} else {
			fmt.Println("Pilihan tidak valid. Ketik 'hit' atau 'stand'")
		}
	}
}

func giliranBandar(tangan *[maksKartu]string, jumlah *int) {
	var sliceTangan []string
	sliceTangan = tangan[:*jumlah]
	var tanganArr [maksKartu]string
	copy(tanganArr[:], sliceTangan)

	fmt.Println("Kartu Bandar:", sliceTangan, "Total:", hitungNilaiTangan(tanganArr, *jumlah))
	for hitungNilaiTangan(tanganArr, *jumlah) < 16 {
		if *jumlah >= maksKartu {
			break
		}
		tangan[*jumlah] = ambilKartu()
		*jumlah = *jumlah + 1

		sliceTangan = tangan[:*jumlah]
		copy(tanganArr[:], sliceTangan)

		fmt.Println("Bandar menarik kartu:", sliceTangan, "Total:", hitungNilaiTangan(tanganArr, *jumlah))
	}
}

func tentukanPemenang(tanganPemain [maksKartu]string, jumlahPemain int, tanganBandar [maksKartu]string, jumlahBandar int) string {
	var totalPemain int
	var totalBandar int
	var hasil string

	totalPemain = hitungNilaiTangan(tanganPemain, jumlahPemain)
	totalBandar = hitungNilaiTangan(tanganBandar, jumlahBandar)

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("                   HASIL PERMAINAN")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("Total Anda: %d | Total Bandar: %d\n", totalPemain, totalBandar)

	if totalPemain > 21 {
		fmt.Println("Anda kalah!")
		fmt.Println("Ingat, 90% penjudi berhenti sebelum menang besar")
		fmt.Println("Ayo pertaruhkan seluruh keluarga anda sekarang!!")
		hasil = "kalah"
	} else if totalBandar > 21 || totalPemain > totalBandar {
		fmt.Println("Anda menang!")
		hasil = "menang"
	} else if totalPemain < totalBandar {
		fmt.Println("Bandar menang!")
		fmt.Println("Ingat, 90% penjudi berhenti sebelum menang besar")
		fmt.Println("Ayo pertaruhkan seluruh keluarga anda sekarang!!")
		hasil = "kalah"
	} else {
		fmt.Println("Seri!")
		hasil = "seri"
	}
	fmt.Println(strings.Repeat("=", 50))

	return hasil
}

func main() {
	var stat Statistik
	var pilihan int
	var tanganPemain [maksKartu]string
	var tanganBandar [maksKartu]string
	var jumlahPemain int
	var jumlahBandar int
	var hasil string

	rand.Seed(time.Now().UnixNano())

	stat.Menang = 0
	stat.Kalah = 0
	stat.Seri = 0

	for {
		tampilkanMenu(stat)
		fmt.Scanln(&pilihan)

		if pilihan == 1 {
			tanganPemain, jumlahPemain, tanganBandar, jumlahBandar = mulaiPermainan()

			var slicePemain []string
			var sliceBandar []string
			slicePemain = tanganPemain[:jumlahPemain]
			sliceBandar = tanganBandar[:1]

			tampilkanKartu(slicePemain, sliceBandar, false)

			giliranPemain(&tanganPemain, &jumlahPemain)

			sliceBandar = tanganBandar[:jumlahBandar]
			tampilkanKartu(tanganPemain[:jumlahPemain], sliceBandar, true)

			giliranBandar(&tanganBandar, &jumlahBandar)

			hasil = tentukanPemenang(tanganPemain, jumlahPemain, tanganBandar, jumlahBandar)

			if hasil == "menang" {
				stat.Menang = stat.Menang + 1
			} else if hasil == "kalah" {
				stat.Kalah = stat.Kalah + 1
			} else if hasil == "seri" {
				stat.Seri = stat.Seri + 1
			}

			fmt.Println("Tekan Enter untuk melanjutkan...")
			var input string
			fmt.Scanln(&input)

		} else if pilihan == 2 {
			fmt.Println("Terima kasih telah bermain!")
			break
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
