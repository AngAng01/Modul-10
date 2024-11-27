package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func hitungRerata(arrBerat arrBalita, n int) float64 {
	var total float64
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	var berat arrBalita
	var bMin, bMax float64

	fmt.Print("Masukkan banyak data berat balita : ")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d : ", i+1)
		fmt.Scanln(&berat[i])
	}

	hitungMinMax(berat, n, &bMin, &bMax)
	rerata := hitungRerata(berat, n)

	fmt.Printf("\nBerat balita minimum : %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum : %.2f kg\n", bMax)
	fmt.Printf("Berat balita rata-rata : %.2f kg\n", rerata)
}
