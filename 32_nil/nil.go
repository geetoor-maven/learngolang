package main

import "fmt"

func main(){
	/*
		NIL GOLANG
		Nil di golang merupakan data yang kosong ( tidak ada value atau nilai di dalamnya ), seperti di bahasa pemerograman lain
		jika suatu object yang belum di assign nilai di dalamnya akan secara otomatis bernilai null atau nil.

		tapi di bahasa pemerograman golang berbeda, nilai nil hanya dapat di gunakan pada tipe data tertentu, seperti slice, function, map, pointer dan channel.
	*/
	

	// contoh inisialisasi data map dengan nilai nil
	var profile map[string] string = nil;
	fmt.Println(profile)

	// contoh inisialisasi data map dengan nilai nil dan validasi jika datanya nil maka berikan nilai di dalam nya
	var animal map[string] string = nil

	if animal == nil {
		animal = map[string] string{
			"name" : "Cow",
		}
	}

	fmt.Println(animal)

	// contoh inisialisasi data nil di slice
	myNames := make([]int, 0)
	fmt.Println(myNames)

	sliceNotNill := BuildSlice(2)
	fmt.Println(sliceNotNill)

	angka := NewMap(0)
	fmt.Println(angka)
}

func NewMap(value int) map[int]int{
	if value == 0 {
		return nil
	}

	return map[int]int{
		1 : 1,
	}
}

func BuildSlice(size int)[]int{
	if size == 0 {
		return nil
	}
	return make([]int, size)
}