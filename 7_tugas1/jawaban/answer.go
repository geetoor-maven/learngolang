package main

import "fmt"

func main(){
	/*
		JANGAN LANGSUNG KE SINI KALAU BELUM CARI TAU ATAU BERFIKIR CARA PENYELESAIAN TUGASNYA HEHE.



		BERIKUT JAWABAN.
	*/

	for{
		fmt.Println("------------------")
		fmt.Println("PROGRAM KALKULATOR SEDERHANA DI GOLANG.");

		var value1, value2, result int;
	
		fmt.Println("1. OPERASI PERKALIAN")
		fmt.Println("2. OPERASI PEMBAGIAN")
		fmt.Println("3. OPERASI PENAMBAHAN")
		fmt.Println("4. OPERASI PENGURANGAN")
		fmt.Println("5. OPERASI MODULUS")
		fmt.Println("6. EXIT")
		fmt.Println("------------------")
		fmt.Print("Pilih salah satu angka di atas : ")
	
		var scan int8;
		fmt.Scanln(&scan);

		if scan == 6 {
			break;
		}

		switch{
		case scan == 1:
			fmt.Println("------------------")
			fmt.Println("OPERASI PERKALIAN")
			fmt.Print("- Masukan Angka Pertama : ")
			fmt.Scanln(&value1);
		
			fmt.Print("- Masukan Angka Kedua : ")
			fmt.Scanln(&value2);
		
			result = value1 * value2;
			fmt.Println("- HASIL DARI PERKALIAN : ", result);
		case scan == 2:
			fmt.Println("------------------")
			fmt.Println("OPERASI PEMBAGIAN")
			fmt.Print("- Masukan Angka Pertama : ")
			fmt.Scanln(&value1);
		
			fmt.Print("- Masukan Angka Kedua : ")
			fmt.Scanln(&value2);
		
			result = value1 / value2;
			fmt.Println("- HASIL DARI PEMBAGIAN : ", result);
		case scan == 3:
			fmt.Println("------------------")
			fmt.Println("OPERASI PENAMBAHAN")
			fmt.Print("- Masukan Angka Pertama : ")
			fmt.Scanln(&value1);
		
			fmt.Print("- Masukan Angka Kedua : ")
			fmt.Scanln(&value2);
		
			result = value1 + value2;
			fmt.Println("- HASIL DARI PENAMBAHAN : ", result);
		case scan == 4:
			fmt.Println("------------------")
			fmt.Println("OPERASI PENGURANGAN")
			fmt.Print("- Masukan Angka Pertama : ")
			fmt.Scanln(&value1);
		
			fmt.Print("- Masukan Angka Kedua : ")
			fmt.Scanln(&value2);
		
			result = value1 - value2;
			fmt.Println("- HASIL DARI PENGURANGAN : ", result);
		case scan == 5:
			fmt.Println("------------------")
			fmt.Println("OPERASI MODULUS")
			fmt.Print("- Masukan Angka Pertama : ")
			fmt.Scanln(&value1);
		
			fmt.Print("- Masukan Angka Kedua : ")
			fmt.Scanln(&value2);
		
			result = value1 % value2;
			fmt.Println("- HASIL DARI MODULUS : ", result);
		default :
			{
				fmt.Println();
				fmt.Println("PILIH ANGKA YANG SESUAI.")
				fmt.Println();
			}	
		}


	}




}