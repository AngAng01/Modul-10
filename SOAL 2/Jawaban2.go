package main

import "fmt"

type IkanWeights [1000]float64

func hitungTotalPerWadah(weights IkanWeights, x, y int) []float64 {
	var totalBerat []float64
	var wadahBerat float64
	for i := 0; i < x; i++ {
		wadahBerat += weights[i]
		if (i+1)%y == 0 || i == x-1 {
			totalBerat = append(totalBerat, wadahBerat)
			wadahBerat = 0
		}
	}
	return totalBerat
}

func hitungRataRata(totalBerat []float64) float64 {
	var total float64
	for i := 0; i < len(totalBerat); i++ {
		total += totalBerat[i]
	}
	return total / float64(len(totalBerat))
}

func main() {
	const maxIkan = 1000
	var x, y int

	fmt.Print("Masukkan jumlah ikan (x) dan kapasitas wadah (y) : ")
	fmt.Scan(&x, &y)

	if x <= 0 || y <= 0 || x > maxIkan {
		fmt.Printf("Input tidak valid.")
		return
	}

	var weights IkanWeights

	fmt.Println("Masukkan berat ikan :")
	for i := 0; i < x; i++ {
		fmt.Printf("Ikan ke %d : ", i+1)
		fmt.Scan(&weights[i])
	}

	totalBerat := hitungTotalPerWadah(weights, x, y)
	rataRata := hitungRataRata(totalBerat)

	fmt.Println("\nTotal berat ikan di setiap wadah :")
	for i := 0; i < len(totalBerat); i++ {
		fmt.Printf("Wadah %d : %.2f\n", i+1, totalBerat[i])
	}

	fmt.Printf("\nBerat rata-rata ikan di setiap wadah : %.2f\n", rataRata)
}
