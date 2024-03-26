package main

import "fmt"

// function yang berfungsi untuk men total score yang berada pada slices
func calculateAverage(scores []int) float64 {

	//menampung angka total score
	sum := 0

	/*menghitung jumlah total angka score di dalam slice kemudian
	dijumlahkan dengan variabel sum untuk menghitung total nilai dari slice score
	*/
	for _, score := range scores {
		sum += score
	}

	/*
		menghitung rata rata dengan mengembalikan nilai dari variabel sum ke float64
		dan membaginya dengan panjang atau total dari slices scores yang merupakan parameter
		dari function totalAverage
	*/
	return float64(sum) / float64(len(scores))

}

// function yang berfungsi untuk menghitung total dari data bonus dan bonus 2
func sumBonus(scores []int) int8 {
	//menampung angka total score
	sum := 0

	//menghitung jumlah total angka dalam slice
	for _, score := range scores {
		sum += score
	}

	/*
		menghitung total nilai dengan mengembalikan nilai dari variabel sum ke int
		dan menambahnya dengan panjang atau total dari slices scores yang merupakan parameter
		dari function sumBonus1
	*/
	return int8(sum) + int8(len(scores))
}

func main() {
	//Skor pada Data 1
	timLumbaLumba := []int{96, 108, 89}
	timKoala := []int{88, 91, 110}

	//menghitung rata ratanya dengan memanggil function calculateAverage
	avgLumbaLumba := calculateAverage(timLumbaLumba)
	avgKoala := calculateAverage(timKoala)

	//mencetak ke konsol untuk rata ratanya
	fmt.Printf("Angka Rata-Rata Lumba-Lumba %.2f\n", avgLumbaLumba)
	fmt.Printf("Angka Rata-Rata Koala %.2f\n", avgKoala)

	//Mengecek pemenangnya, jika memenuhi persyaratan yang pertama maka menang,
	//jika tidak maka kalah dan jika tidak memenuhi keduanya maka hasilnya seri
	switch {
	case avgLumbaLumba > avgKoala:
		fmt.Println("Pemenang Data 1 : Tim Lumba Lumba")
	case avgKoala > avgLumbaLumba:
		fmt.Println("Pemenang Data 1 : Tim Koala")
	default:
		fmt.Println("Seri")
	}

	//Data Bonus1
	lumbaLumba := []int{97, 112, 101}
	koala := []int{109, 95, 123}

	//Menampung data yang sudah diatas 100
	koalaAbove100 := []int{}
	lumbaLumbaAbove100 := []int{}

	for _, score := range lumbaLumba {
		if score > 100 {
			/*
				setelah dicek kemudian menggunakan method append untuk
				menambahkan score ke dalam variabel lumbaLumbaAbove100
			*/
			lumbaLumbaAbove100 = append(lumbaLumbaAbove100, score)
			continue
		}
	}

	for _, score := range koala {
		if score > 100 {
			/*
				setelah dicek kemudian menggunakan method append untuk
				menambahkan score ke dalam variabel lumbaLumbaAbove100
			*/
			koalaAbove100 = append(koalaAbove100, score)
			continue
		}
	}

	//menjumlah skor yang sudah diatas 100 dengan memanggil function sumBonus
	totalKoalaAbove100 := sumBonus(koalaAbove100)
	totalLumbaAbove100 := sumBonus(lumbaLumbaAbove100)

	switch {
	case totalKoalaAbove100 > totalLumbaAbove100:
		fmt.Println("Pemenang Bonus 1 : Koala")
	case totalLumbaAbove100 > totalKoalaAbove100:
		fmt.Println("Pemenang Bonus 1 : Lumba Lumba")
	default:
		fmt.Println("Seri")
	}

	//Data Bonus2
	lumba2 := []int{97, 112, 101}
	koala2 := []int{109, 95, 106}

	totalLumba2 := sumBonus(lumba2)
	totalKoala2 := sumBonus(koala2)

	switch {
	case totalLumba2 > totalKoala2:
		fmt.Println("Pemenang Bonus 2 : Lumba Lumba")
	case totalKoala2 > totalLumba2:
		fmt.Println("Pemenang Bonus 2 : Koala")
	default:
		fmt.Println("Pemenang Bonus 2 : Seri, Tidak ada pemenang")
	}
}
