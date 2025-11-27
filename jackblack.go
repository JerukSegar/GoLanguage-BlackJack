package main

import (
	"fmt"
	"math/rand"
	"time"
)

const maksKartu = 10 

var kartu = [13]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
var nilaiKartu = map[string]int{"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10, "J": 10, "Q": 10, "K": 10, "A": 11}

func ambilKartu() string {
	return kartu[rand.Intn(len(kartu))]
}

func hitungNilaiTangan(tangan [maksKartu]string, jumlah int) int {
	total, jumlahAs := 0, 0

	for i := 0; i < jumlah; i++ {
		k := tangan[i]
		total += nilaiKartu[k]
		if k == "A" {
			jumlahAs++
		}
	}

	for total > 21 && jumlahAs > 0 {
		total -= 10
		jumlahAs--
	}

	return total
}

func mulaiPermainan() ([maksKartu]string, int, [maksKartu]string, int) {
	var tanganPemain, tanganBandar [maksKartu]string
	tanganPemain[0], tanganPemain[1] = ambilKartu(), ambilKartu()
	tanganBandar[0], tanganBandar[1] = ambilKartu(), ambilKartu()

	fmt.Println("Kartu Anda:", tanganPemain[:2], "Total:", hitungNilaiTangan(tanganPemain, 2))
	return tanganPemain, 2, tanganBandar, 2
}

func giliranPemain(tangan *[maksKartu]string, jumlah *int) {
	for {
		fmt.Print("Hit / Stand")
		var pilihan string
		fmt.Scanln(&pilihan)

		if pilihan == "hit" {
			if *jumlah >= maksKartu {
				fmt.Println("Anda sudah memiliki terlalu banyak kartu!")
				break
			}
			tangan[*jumlah] = ambilKartu()
			*jumlah++
			fmt.Println("Kartu Anda:", tangan[:*jumlah], "Total:", hitungNilaiTangan(*tangan, *jumlah))
			if hitungNilaiTangan(*tangan, *jumlah) > 21 {
				fmt.Println("Bust!")
				return
			}
		} else {
			break
		}
	}
}

func giliranBandar(tangan *[maksKartu]string, jumlah *int) {
	fmt.Println("Kartu Bandar:", tangan[:*jumlah], "Total:", hitungNilaiTangan(*tangan, *jumlah))
	for hitungNilaiTangan(*tangan, *jumlah) < 16 {
		if *jumlah >= maksKartu {
			break
		}
		tangan[*jumlah] = ambilKartu()
		*jumlah++
		fmt.Println("Bandar menarik kartu:", tangan[:*jumlah], "Total:", hitungNilaiTangan(*tangan, *jumlah))
	}
}

func tentukanPemenang(tanganPemain [maksKartu]string, jumlahPemain int, tanganBandar [maksKartu]string, jumlahBandar int) {
	totalPemain := hitungNilaiTangan(tanganPemain, jumlahPemain)
	totalBandar := hitungNilaiTangan(tanganBandar, jumlahBandar)

	if totalPemain > 21 {
		fmt.Println("Anda kalah!")
		fmt.Println("ingat, 90% penjudi berhenti sebelum menang besar")
		fmt.Println("Ayo pertaruhkan seluruh keluarga anda sekarang!!")
	} else if totalBandar > 21 || totalPemain > totalBandar {
		fmt.Println("Anda menang!")
	} else if totalPemain < totalBandar {
		fmt.Println("Bandar menang!")
		fmt.Println("ingat, 90% penjudi berhenti sebelum menang besar")
		fmt.Println("Ayo pertaruhkan seluruh keluarga anda sekarang!!")
	} else {
		fmt.Println("Seri!")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	tanganPemain, jumlahPemain, tanganBandar, jumlahBandar := mulaiPermainan()
	giliranPemain(&tanganPemain, &jumlahPemain)
	giliranBandar(&tanganBandar, &jumlahBandar)
	tentukanPemenang(tanganPemain, jumlahPemain, tanganBandar, jumlahBandar)
}
