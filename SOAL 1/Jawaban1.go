package main

import "fmt"

type arrBerat [1000]float64

func BeratTerkecil(arr arrBerat, n int) float64 {
	min := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

func BeratTerbesar(arr arrBerat, n int) float64 {
	max := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func main() {
	var N int
	var berat arrBerat

	fmt.Print("Masukkan jumlah anak kelinci (N) : ")
	fmt.Scan(&N)

	fmt.Printf("Masukkan berat %d anak kelinci :\n", N)
	for i := 1; i <= N; i++ {
		fmt.Printf("Kelinci ke %d : ", i)
		fmt.Scan(&berat[i])
	}

	min := BeratTerkecil(berat, N)
	max := BeratTerbesar(berat, N)

	fmt.Printf("Berat terkecil : %.2f\n", min)
	fmt.Printf("Berat terbesar : %.2f\n", max)
}
