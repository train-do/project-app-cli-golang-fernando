package views

import (
	"app/models"
	"fmt"
)

func MainMenuView() {
	fmt.Println("1. Login\n2. Registrasi Akun\n3. Tampilkan All Data User\n0. Exit")
	fmt.Print("Masukkan menu : ")
}
func PrintAllUser(data []models.User)  {
	for _, v := range data {
		fmt.Print("Username : ", v.Username)
		fmt.Print(" -- ")
		fmt.Print("Nama : ", models.PublicInformasiName(v.Person))
		fmt.Print(" -- ")
		fmt.Print("Bank : ", models.PublicInformasiName(v.Bank))
		fmt.Print(" -- ")
		fmt.Println("Saldo : ", v.Bank.Saldo)
	}
}
func PrintMBanking()  {
	fmt.Println("1. Cek Saldo")
	fmt.Println("2. Simpan Tunai")
	fmt.Println("3. Tarik Tunai")
	fmt.Println("4. History Transaksi")
	fmt.Println("0. Logout ")
}
func PrintHistory(history []models.History)  {
	if len(history) == 0 {
		fmt.Println("Belum Ada Transaksi")
	}else{
		for _, v := range history {
			fmt.Printf("%+v\n", v)
		}
	}
}