package main

import (
	"fmt"
)

const passDb string = "123456";
var saldo int = 0;

func main(){
	/*
		JANGAN LANGSUNG KE SINI KALAU BELUM CARI TAU ATAU BERFIKIR CARA PENYELESAIAN TUGASNYA HEHE.



		BERIKUT JAWABAN Tugas 4.
	*/
	line()
	fmt.Println("WELCOME TO CLI BANK")
	line()
	viewApps()
}

func viewApps(){
	for{
		fmt.Println("1. Masuk Applikasi")
		fmt.Println("2. Keluar Applikasi")

		var inputAngka = input("Masukan angka 1/2 :");
		
		if inputAngka == "1" {
			loginView()
			break;
		}else if inputAngka == "2"{
			line()
			outPut("Berhasil Keluar applikasi..")
			line()
			break;
		}else {
			outPut("Masukan Angka yang benar...")
		}
	}
}

func loginView(){
	for{

		var passInput = input("Masukan password :");
		isValidate := validate
		if validateLogin(passInput, isValidate) {
			break
		}

	}
}

func validateLogin(input string, validation func(string) bool)bool{
	return validation(input)
}

func validate(password string)bool{
	isValidate := true;
	if  len(password) < 6 {
		line()
		outPut("Password tidak boleh kurang dari 6 angka")
		line()
		isValidate = false;
	}else if len(password) > 6 {
		line()
		outPut("Password tidak boleh lebih dari 6 angka")
		line()
		isValidate = false;
	}else if password == passDb {
		line()
		outPut("Login Berhasil.....")
		line()
		viewDashboard()
		isValidate = true;
	}else{
		line()
		outPut("Password anda salah.....")
		line()
		isValidate = false;
	}
	return isValidate;
}

func viewDashboard(){
	for{
		variadicFeatureDashboard("Lihat Saldo", "Top Up Saldo", "Kirim Uang", "Tarik Uang", "Lihat History Transaksi", "Exit")
		inputFeature := input("Masukan pilihan Feature 1-6 : ")
		if inputFeature == "6" {
			viewApps()
			break
		}

		if inputFeature == "1" {
			line();
			viewSaldo();
			line()
		}else if inputFeature == "2"{
			line()
			topUpSaldo()
			line();
		}
	}
}

func variadicFeatureDashboard(features ...string){
	for i, feature := range features{
		fmt.Println((i+1), feature)
	}
}

func topUpSaldo(){
	fmt.Print("Masukan total uang yang ingin di top up : ")
	var inputSaldo int;
	fmt.Scanln(&inputSaldo)
	if inputSaldo == 0 {
		outPut("Saldo yang di masukan tidak boleh kosong")
	}else{
		saldo = inputSaldo;
		outPut("Berhasil TopUp.....")
	}
}

func viewSaldo(){
	fmt.Println("Total saldo anda :",saldo)
}

func line(){
	fmt.Println("--------------------");
}

func outPut(value string){
	fmt.Println(value)
}

func input(message string)string{
	fmt.Print(message)
	var result string;
	fmt.Scanln(&result);
	return result ;
}
