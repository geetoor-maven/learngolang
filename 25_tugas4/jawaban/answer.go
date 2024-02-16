package main

import (
	"fmt"
	"strconv"
)

const passDb string = "123456";
var saldo int = 0;
var historyTransactions = []string{}

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
	loginView := func ()  {
		for{
			var passInput = input("Masukan password :");
			isValidate := validate
			if validateLogin(passInput, isValidate) {
				break
			}
		}
	}
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



func validateLogin(input string, validation func(string) bool)bool{
	return validation(input)
}

func validate(password string)bool{
	isValidate := false;
	if  len(password) < 6 {
		line()
		outPut("Password tidak boleh kurang dari 6 angka")
		line()
	}else if len(password) > 6 {
		line()
		outPut("Password tidak boleh lebih dari 6 angka")
		line()
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
			outPut("TopUp Saldo....")
			topUpSaldo()
			line();
		}else if inputFeature == "3"{
			line()
			outPut("Kirim Uang....")
			kirimSaldo()
			line()
		}else if inputFeature == "4"{
			line()
			outPut("Tarik Uang....")
			tarikSaldo()
			line()
		}else if inputFeature == "5"{
			line()
			outPut("History Transaksi....")
			viewHistory()
			line()
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
		saldo = saldo + inputSaldo;
		outPut("Berhasil TopUp.....")
		history := "TopUp Saldo " + strconv.Itoa(inputSaldo)
		setDataToHistory(history)
	}
}

func tarikSaldo(){
	fmt.Print("Masukan total uang yang ingin di Tarik : ")
	var totalTarik int;
	fmt.Scanln(&totalTarik)
	validate, _ , _ := validasiTransaksi("4", 0, totalTarik);
	if validate {
		saldo = saldo - totalTarik;
		outPut("Berhasil Tarik Uang uang...")
		history := "Tarik Saldo Sebesar " + strconv.Itoa(totalTarik)
		setDataToHistory(history)
	}else{
		outPut("Total uang yang anda ingin di tarik melebihi dari saldo anda.")
	}
}

func kirimSaldo(){
	var rekeningTujuan = input("Masukan nomor rekening tujuan : ")
	if len(rekeningTujuan) == 5 {
	
		fmt.Print("Masukan total uang yang ingin di kirim : ")
		var totalKirim int;
		fmt.Scanln(&totalKirim)

		validate, _ , _ := validasiTransaksi("3", totalKirim, 0);
		if validate {
			saldo = saldo - totalKirim;
			outPut("Berhasil kirim uang...")
			history := "Kirim Saldo ke " + rekeningTujuan + " Sebesar " + strconv.Itoa(totalKirim)
			setDataToHistory(history)
		}else{
			outPut("Total uang yang anda kirimkan melebihi dari saldo anda.")
		}

	}else {
		outPut("Rekening Tujuan tidak boleh lebih dan kurang dari 5")
	}
}

func validasiTransaksi(code string, totalKirim int, totalTarik int)(bool, int, int){
	isValidate := false;
	if code == "3" {
		if totalKirim < saldo {
			isValidate = true;
		}else {
			isValidate = false;
		}
	}else if code == "4"{
		if totalTarik < saldo {
			isValidate = true;
		}else {
			isValidate = false;
		}
	}
	return isValidate, totalKirim, totalTarik;
}


func setDataToHistory(message string){
	historyTransactions = append(historyTransactions, message)
}

func viewHistory(){
	for i := 0; i < len(historyTransactions); i++ {
		if historyTransactions[i] != "" {
			fmt.Println("- ", ":", historyTransactions[i]);
		}
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
