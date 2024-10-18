package controllers

import (
	"app/models"
	"errors"
	"fmt"
	"strconv"
)

func ValidasiMenu(input string, opsiMenu int) error {
	isValid := false
	for i := 0; i <= opsiMenu; i++ {
		if input == strconv.Itoa(i) {
			isValid = true
		}
	}
	if !isValid {
		str := "Inputan hanya diisi angka 0 dan angka 1 sampai "+ fmt.Sprintf("%d", opsiMenu)
		return errors.New(str)
	}
	return nil
}
func ValidasiIsNumber(input string) error {
    _, err := strconv.Atoi(input)
	if err != nil {
		return errors.New("Inputan bukan angka")
	}
	return nil
}
func ValidasiIsAlphabet(input string) error {
    _, err := strconv.Atoi(input)
	if err == nil {
		return errors.New("Inputan bukan alphabet")
	}
	return nil
}
func ValidasiNegativeNumber(input string) error {
	integer, _ := strconv.Atoi(input)
	if integer <= 0 {
		return errors.New("Inputan harus positif dan tidak boleh 0")
	}
	return nil
}
func ValidasiLogin(input string, database *[]models.User) (int, error) {
	for i, v := range *database {
		if v.Username ==  input {
			return i, nil
		}
	}
	return -1, errors.New("Salah Username")
}
func ValidasiLimitSaldo(input string, saldo int) error {
    withdraw, _ := strconv.Atoi(input)
	if withdraw > saldo {
		return errors.New("Saldo Tidak Mencukupi")
	}
	return nil
}