package controllers

import (
	"app/models"
	"app/views"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)
var history []models.History
func RegisterUser(database *[]models.User) {
	form := []string{"Username : ", "Nama : ", "Alamat : ", "Bank : ", "Saldo : "}
	var property []string
	for i := 0; i < len(form); i++ {
		var input string
		var isNumber, isAlphabet, isPositiveNumber error
		fmt.Print(form[i])
		fmt.Scanln(&input)
		if i == len(form)-1 {
			isNumber = ValidasiIsNumber(input)
			isPositiveNumber = ValidasiNegativeNumber(input)
		}else{
			isAlphabet = ValidasiIsAlphabet(input)
		}
		if isNumber != nil {
			ClearScreen()
			fmt.Println(isNumber)
			i--
			continue
		}
		if isPositiveNumber != nil {
			ClearScreen()
			fmt.Println(isPositiveNumber)
			i--
			continue
		}
		if isAlphabet != nil {
			ClearScreen()
			fmt.Println(isAlphabet)
			i--
			continue
		}
		property = append(property, input)
		ClearScreen()
	}
	saldo, _ := strconv.Atoi(property[4])
	newUser := models.CreateUser(property[0], property[1], property[2], property[3], saldo)
	*database = append(*database, newUser)
	ClearScreen()
	fmt.Println("Sukses Registrasi")
}
func Login(database *[]models.User) {
	for{
		var input string
		fmt.Print("Masukkan Username : ")
		fmt.Scanln(&input)
		indeks, isValid := ValidasiLogin(input, database)
		if isValid != nil{
			ClearScreen()
			fmt.Println(isValid)
			continue
		}
		mBanking(&(*database)[indeks])
		break
	}
	ClearScreen()
}
func mBanking(dataUser *models.User){
	ClearScreen()
	for{
		views.PrintMBanking()
        var input string
		fmt.Print("Masukkan menu : ")
        fmt.Scanln(&input)
        validasi := ValidasiMenu(input, 4)
        if validasi != nil {
			ClearScreen()
            fmt.Println(validasi)
            continue
        }
		switch input {
		case "1":
			ClearScreen()
			fmt.Printf("Saldo anda tersisa %d rupiah\n", dataUser.Bank.Saldo)
		case "2":
			ClearScreen()
			DepositCash(dataUser)
		case "3":
			ClearScreen()
			WithdrawCash(dataUser)
		case "4":
			getHistory(dataUser.Username)
		case "0":
			ClearScreen()
			return
		}
	}
}

func DepositCash(dataUser *models.User)  {
	for{
		var saldo string
		fmt.Print("Masukkan nominal deposit: ")
		fmt.Scanln(&saldo)
		var isNumber, isPositiveNumber error
		isNumber = ValidasiIsNumber(saldo)
		isPositiveNumber = ValidasiNegativeNumber(saldo)
		if isNumber != nil {
			ClearScreen()
			fmt.Println(isNumber)
			continue
		}
		if isPositiveNumber != nil {
			ClearScreen()
			fmt.Println(isPositiveNumber)
			continue
		}
		value, _ := strconv.Atoi(saldo)
		dataUser.Bank.Saldo += value
		history = append(history, models.History{dataUser.Username, "Deposit", value})
		break
	}
	ClearScreen()
}
func WithdrawCash(dataUser *models.User)  {
	for{
		var saldo string
		fmt.Print("Masukkan nominal penarikan: ")
		fmt.Scanln(&saldo)
		var isNumber, isPositiveNumber, isLimit error
		isNumber = ValidasiIsNumber(saldo)
		isPositiveNumber = ValidasiNegativeNumber(saldo)
		isLimit = ValidasiLimitSaldo(saldo, dataUser.Bank.Saldo)
		if isNumber != nil {
			ClearScreen()
			fmt.Println(isNumber)
			continue
		}
		if isPositiveNumber != nil {
			ClearScreen()
			fmt.Println(isPositiveNumber)
			continue
		}
		if isLimit != nil {
			ClearScreen()
			fmt.Println(isLimit)
			fmt.Println("Balance : ", dataUser.Bank.Saldo)
			continue
		}
		value, _ := strconv.Atoi(saldo)
		dataUser.Bank.Saldo -= value
		history = append(history, models.History{dataUser.Username, "Withdraw", value})
		break
	}
	ClearScreen()
}
func getHistory(username string)  {
	var userHistory []models.History
	for _, v := range history {
		if v.Username == username {
			userHistory = append(userHistory, v)
		}
	}
	ClearScreen()
	views.PrintHistory(userHistory)
}
func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}