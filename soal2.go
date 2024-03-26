package main

import (
	"fmt"
	"math"
)

// data adalah struct yang berguna untuk menyimpan berat dan tinggi seseorang.
type Data struct {
	Berat  float64 // Berat badan dalam kg
	Tinggi float64 // Tinggi badan dalam meter
}

// hitungBMI adalah fungsi untuk menghitung BMI seseorang yang menggunakan package Math.pow yang berarti pemangkatan
func hitungBMI(berat, tinggi float64) float64 {
	return berat / math.Pow(tinggi, 2)
}

// cetakBMI adalah fungsi untuk mencetak hasil perbandingan BMI. mengambil 2 angka dibelakang koma
func cetakBMI(person string, bmi float64) {
	fmt.Printf("BMI %s: %.2f\n", person, bmi)
}

// fungsi berikut adalah fungsi untuk mencetak hasil perbandingan BMI antara dua orang.
func cetakHasil(person1, person2 string, bmi1, bmi2 float64) {
	fmt.Printf("%s memiliki BMI lebih tinggi dari %s: %t\n", person1, person2, bmi1 > bmi2)
}

func main() {
	// Data
	mark1 := Data{Berat: 78.0, Tinggi: 1.69}
	john1 := Data{Berat: 92.0, Tinggi: 1.95}

	mark2 := Data{Berat: 95.0, Tinggi: 1.88}
	john2 := Data{Berat: 85.0, Tinggi: 1.76}

	// Hitung BMI mereka menggunakan rumus
	bmiMark1 := hitungBMI(mark1.Berat, mark1.Tinggi)
	bmiJohn1 := hitungBMI(john1.Berat, john1.Tinggi)

	bmiMark2 := hitungBMI(mark2.Berat, mark2.Tinggi)
	bmiJohn2 := hitungBMI(john2.Berat, john2.Tinggi)

	// Cetak hasil dengan membatasi jumlah angka dibelakang koma
	fmt.Println("Data 1:")
	cetakBMI("Mark", bmiMark1)
	cetakBMI("John", bmiJohn1)
	cetakHasil("Mark", "John", bmiMark1, bmiJohn1)

	fmt.Println("\nData 2:")
	cetakBMI("Mark", bmiMark2)
	cetakBMI("John", bmiJohn2)
	cetakHasil("Mark", "John", bmiMark2, bmiJohn2)
}
