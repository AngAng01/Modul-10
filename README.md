# MODUL SEPULUH
  ## SOAL 1
  Program yang Anda berikan adalah program yang digunakan untuk menentukan berat terkecil dan terbesar dari sejumlah anak kelinci yang akan dijual berdasarkan input dari pengguna.
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt' dan 'math').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Pendefinisian Tipe Data arrBerat, Sebuah array dengan kapasitas 1000 elemen bertipe float64 yang digunakan untuk menyimpan berat anak kelinci.
      - Fungsi BeratTerkecil, Fungsi ini menerima parameter berupa array arrBerat dan jumlah elemen n dan juga Mengembalikan nilai berat terkecil dari elemen-elemen dalam array.
      - Fungsi BeratTerbesar, Fungsi ini mirip dengan BeratTerkecil, tetapi bertujuan untuk mencari dan mengembalikan berat terbesar dalam array.
      
   ## Code Explanation
   ```go
   type arrBerat [1000]float64
   ```
   Kode di atas adalah deklarasi tipe data baru berupa array dengan kapasitas 1000 elemen bertipe float64.
  
   ```go
  func BeratTerkecil(arr arrBerat, n int) float64 {
  	min := arr[0]
  	for i := 1; i < n; i++ {
  		if arr[i] < min {
  			min = arr[i]
  		}
  	}
  	return min
  }
   ```
   Kode di atas adalah fungsi BeratTerkecil yang mencari nilai terkecil dalam array arrBerat. Fungsi ini menerima dua parameter, yaitu array arr dan integer n, yang menunjukkan jumlah elemen. Fungsi membandingkan elemen-elemen array dan mengembalikan nilai terkecil yang ditemukan.

   ```go
  func BeratTerbesar(arr arrBerat, n int) float64 {
  	max := arr[0]
  	for i := 1; i < n; i++ {
  		if arr[i] > max {
  			max = arr[i]
  		}
  	}
  	return max
  }
   ```
   Kode di atas adalah fungsi BeratTerbesar yang digunakan untuk mencari nilai terbesar dalam array arrBerat. Fungsi ini menerima dua parameter, yaitu array arr dan integer n yang menunjukkan jumlah elemen. Fungsi membandingkan elemen-elemen array dan mengembalikan nilai terbesar yang ditemukan.

   ```go
   var N int
   ```
   Kode diatas adalah deklarasi variabel yang menyatakan bahwa variabel N bertipe int (integer) dan digunakan untuk menyimpan nilai bilangan bulat. 

   ```go
   var berat arrBerat
   ```
   Kode di atas adalah deklarasi variabel yang menyatakan bahwa variabel berat bertipe arrBerat, yaitu array dengan kapasitas 1000 elemen bertipe float64

   ```go
	fmt.Print("Masukkan jumlah anak kelinci (N) : ")
	fmt.Scan(&N)
   ```
   Kode di atas adalah kode untuk memberi pesan dan meminta input dari pengguna dan menyimpannya dalam variabel N.

   ```go
	fmt.Printf("Masukkan berat %d anak kelinci :\n", N)
	for i := 1; i <= N; i++ {
		fmt.Printf("Kelinci ke %d : ", i)
		fmt.Scan(&berat[i])
	}
   ```
   Kode di atas adalah kode untuk meminta pengguna memasukkan berat anak kelinci. Pesan pertama menampilkan jumlah anak kelinci yang harus dimasukkan, dan kemudian menggunakan loop for untuk meminta input berat setiap kelinci satu per satu, menyimpannya dalam array berat.

   ```go
	min := BeratTerkecil(berat, N)
	max := BeratTerbesar(berat, N)
   ```
   Kode di atas memanggil dua fungsi, BeratTerkecil dan BeratTerbesar, untuk mencari nilai terkecil dan terbesar dalam array berat dengan jumlah elemen N. Hasilnya disimpan dalam variabel min dan max yang masing-masing berisi berat terkecil dan terbesar.

   ```go
	fmt.Printf("Berat terkecil : %.2f\n", min)
	fmt.Printf("Berat terbesar : %.2f\n", max)
   ```
   Kode di atas digunakan untuk menampilkan hasil berat terkecil dan terbesar yang sudah dihitung sebelumnya. fmt.Printf("Berat terkecil : %.2f\n", min) mencetak berat terkecil dengan dua angka desimal, sedangkan fmt.Printf("Berat terbesar : %.2f\n", max) mencetak berat terbesar dengan dua angka desimal.


   ## SOAL 2
  Program di atas adalah program yang menghitung total berat ikan di setiap wadah, membagi ikan di beberapa wadah dan menghitung rata-rata berat ikan per wadah.
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt' dan 'math').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Deklarasi Tipe Data IkanWeights, Mendeklarasikan tipe data baru berupa array dengan kapasitas 1000 elemen bertipe float64 untuk menyimpan berat ikan.
      - Fungsi hitungTotalPerWadah, Fungsi untuk menghitung total berat ikan dalam setiap wadah berdasarkan jumlah ikan dan kapasitas wadah yang diberikan.
      - Fungsi hitungRataRata, Fungsi untuk menghitung rata-rata berat ikan di setiap wadah setelah total berat dihitung.
      
   ## Code Explanation
   ```go
   type IkanWeights [1000]float64
   ```
   Kode di atas adalah kode untuk mendeklarasikan tipe data baru yang disebut IkanWeights, yang merupakan array dengan kapasitas 1000 elemen bertipe float64. Tipe data ini digunakan untuk menyimpan berat ikan yang akan dihitung dalam program.
  
   ```go
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
   ```
   Kode di atas adalah fungsi hitungTotalPerWadah yang menghitung total berat ikan dalam setiap wadah. Fungsi ini menjumlahkan berat ikan dan menyimpannya dalam array totalBerat setiap kali kapasitas wadah tercapai atau ketika ikan terakhir diproses.

   ```go
  func hitungRataRata(totalBerat []float64) float64 {
  	var total float64
  	for i := 0; i < len(totalBerat); i++ {
  		total += totalBerat[i]
  	}
  	return total / float64(len(totalBerat))
  }
   ```
  Kode di atas adalah fungsi hitungRataRata yang menghitung rata-rata dari total berat ikan di setiap wadah. Fungsi ini menjumlahkan semua nilai dalam array totalBerat dan membaginya dengan jumlah wadah untuk menghasilkan nilai rata-rata.

   ```go
   const maxIkan = 1000
   ```
   Kode diatas adalah deklarasi konstanta yang menetapkan nilai maksimal jumlah ikan (maxIkan) yang dapat diproses dalam program, yaitu 1000 ikan. 

   ```go
   var x, y int
   ```
   var x, y int adalah deklarasi dua variabel bertipe int (integer). Variabel x digunakan untuk menyimpan jumlah ikan yang akan diproses, sedangkan variabel y menyimpan kapasitas wadah untuk ikan. 

   ```go
	fmt.Print("Masukkan jumlah ikan (x) dan kapasitas wadah (y) : ")
	fmt.Scan(&x, &y)
   ```
   Kode di atas adalah kode fmt.Print("Masukkan jumlah ikan (x) dan kapasitas wadah (y) : ") menampilkan pesan untuk meminta input, sementara fmt.Scan(&x, &y) membaca input jumlah ikan dan kapasitas wadah dari pengguna.

   ```go
	if x <= 0 || y <= 0 || x > maxIkan {
		fmt.Printf("Input tidak valid.")
		return
	}
   ```
   Kode di atas adalah kode untuk memvalidasi input jumlah ikan dan kapasitas wadah, memastikan bahwa jumlah ikan (x) lebih besar dari 0, kapasitas wadah (y) lebih besar dari 0, dan jumlah ikan tidak melebihi batas maksimal (maxIkan).

   ```go
	var weights IkanWeights
   ```
   Kode diatas adalah kode untuk mendeklarasikan variabel weights dengan tipe IkanWeights, yang merupakan array bertipe float64 dengan kapasitas 1000 untuk menyimpan berat ikan yang akan dimasukkan oleh pengguna.

   ```go
	fmt.Println("Masukkan berat ikan :")
	for i := 0; i < x; i++ {
		fmt.Printf("Ikan ke %d : ", i+1)
		fmt.Scan(&weights[i])
	}
   ```
   Kode diatas adalah kode untuk meminta pengguna memasukkan berat ikan satu per satu. Dimulai dengan menampilkan pesan "Masukkan berat ikan:", kemudian melalui loop for, pengguna diminta untuk memasukkan berat ikan ke dalam array weights berdasarkan jumlah ikan (x). Setiap berat ikan dimasukkan ke dalam elemen array yang sesuai dengan urutan ikan.

   ```go
	totalBerat := hitungTotalPerWadah(weights, x, y)
	rataRata := hitungRataRata(totalBerat)
   ```
   Kode diatas adalah kode untuk menghitung total berat ikan di setiap wadah dan berat rata-rata ikan di setiap wadah. Fungsi hitungTotalPerWadah digunakan untuk menghitung total berat ikan dalam setiap wadah berdasarkan array weights, jumlah ikan (x), dan kapasitas wadah (y). Selanjutnya, fungsi hitungRataRata digunakan untuk menghitung rata-rata berat ikan di setiap wadah berdasarkan hasil perhitungan total berat dari setiap wadah.

   ```go
	fmt.Println("\nTotal berat ikan di setiap wadah :")
	for i := 0; i < len(totalBerat); i++ {
		fmt.Printf("Wadah %d : %.2f\n", i+1, totalBerat[i])
	}
   ```
   Kode diatas adalah kode untuk menampilkan total berat ikan di setiap wadah. Menggunakan loop for, program mencetak berat ikan pada setiap wadah berdasarkan hasil dari totalBerat.

   ```go
	fmt.Printf("\nBerat rata-rata ikan di setiap wadah : %.2f\n", rataRata)
   ```
   Kode diatas adalah kode untuk menampilkan berat rata-rata ikan di setiap wadah dengan format dua angka desimal.


   ## SOAL 3
  Kode diatas adalah program yang digunakan untuk menghitung berat balita minimum, maksimum, dan rata-rata berdasarkan input yang diberikan. 
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt' dan 'math').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Deklarasi tipe data arrBalita, Menggunakan array dengan kapasitas 100 untuk menampung data berat balita.
      - Fungsi hitungMinMax, Menghitung nilai berat balita minimum dan maksimum dengan menggunakan pointer untuk mengubah nilai variabel di luar fungsi.
      - Fungsi hitungRerata, Menghitung rata-rata berat balita berdasarkan jumlah data yang dimasukkan.
      
   ## Code Explanation
   ```go
   type arrBalita [100]float64
   ```
   Kode di atas adalah kode untuk deklarasi tipe data baru dengan nama arrBalita, yang merupakan sebuah array dengan kapasitas 100 elemen bertipe float64. 
  
   ```go
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
   ```
   Kode diatas adalah kode fungsi hitungMinMax yang menghitung berat balita minimum dan maksimum. Fungsi ini menerima array berat balita, jumlah data, serta pointer untuk menyimpan hasil minimum dan maksimum. Fungsi ini menginisialisasi nilai bMin dan bMax dengan elemen pertama array, kemudian mencari nilai terkecil dan terbesar dari array.

   ```go
  func hitungRerata(arrBerat arrBalita, n int) float64 {
  	var total float64
  	for i := 0; i < n; i++ {
  		total += arrBerat[i]
  	}
  	return total / float64(n)
  }
   ```
  Kode diatas adalah kode fungsi hitungRerata yang menghitung rata-rata berat balita. Fungsi ini menerima array berat balita dan jumlah data, kemudian menjumlahkan semua elemen dalam array dan membaginya dengan jumlah data untuk mendapatkan rata-rata.
  
   ```go
   var n int
   ```
   Kode diatas adalah kode deklarasi variabel n yang bertipe integer, yang digunakan untuk menyimpan jumlah data berat balita yang akan dimasukkan oleh pengguna.

   ```go
  var berat arrBalita
   ```
   Kode diatas adalah kode deklarasi variabel berat dengan tipe data arrBalita, yang merupakan array dengan kapasitas 100 elemen bertipe float64. Variabel ini digunakan untuk menyimpan data berat balita yang akan dimasukkan oleh pengguna.

   ```go
	var bMin, bMax float64
   ```
   Kode diatas adalah kode deklarasi dua variabel bMin dan bMax dengan tipe data float64, yang digunakan untuk menyimpan nilai berat balita terkecil dan terbesar setelah proses perhitungan dilakukan.

   ```go
	fmt.Print("Masukkan banyak data berat balita : ")
	fmt.Scanln(&n)
   ```
   Kode diatas adalah kode untuk meminta input dari pengguna berupa jumlah data berat balita yang akan dimasukkan, kemudian nilai tersebut disimpan dalam variabel n yang bertipe int.

   ```go
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d : ", i+1)
		fmt.Scanln(&berat[i])
	}
   ```
   Kode diatas adalah kode untuk melakukan perulangan sebanyak n kali, yang masing-masing kali meminta input berat balita ke-i dan menyimpannya ke dalam array berat.

   ```go
	hitungMinMax(berat, n, &bMin, &bMax)
	rerata := hitungRerata(berat, n)
   ```
   Kode diatas adalah kode untuk memanggil fungsi hitungMinMax yang akan mengisi nilai minimum (bMin) dan maksimum (bMax) dari array berat, dan memanggil fungsi hitungRerata untuk menghitung nilai rata-rata berat balita dari array berat.

   ```go
	fmt.Printf("\nBerat balita minimum : %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum : %.2f kg\n", bMax)
	fmt.Printf("Berat balita rata-rata : %.2f kg\n", rerata)
   ```
   Kode diatas adalah kode untuk mencetak hasil perhitungan berat balita, yaitu berat balita minimum, maksimum, dan rata-rata dalam format desimal dua angka setelah koma (kg).
