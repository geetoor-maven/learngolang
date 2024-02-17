package main

import "fmt"

func main(){
	/*
		RECOVER GOLANG.
		function recover di gunakan untuk menangkap error ( data panic function ), sehingga program tidak akan berhenti.

		function recover() di golang sama halnya exception di bahasa pemerograman java yang dapat menangkap error.
	*/

	// recover atau exception tidak di tangkap di panic function karena parameternya false
	run(false)

	// recover atau exception  di tangkap di panic function karena parameternya true
	run(true)
}

func run(err bool){
	defer recoverError()

	if err {
		panic("erorr di tangkap")
	}

	fmt.Println("Run Application...")
	fmt.Println("On Progress.....")
}

func recoverError(){

	// variable err akan menyimpan exception data dari panic function
	err := recover()
	// handle variable err jika tidak nil maka munculkan log error
	if err != nil {
		fmt.Println("log Error :", err)	
	}

	fmt.Println("Done Run....")
	fmt.Println("............")
}