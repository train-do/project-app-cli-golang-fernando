package models

type User struct {
	Username string
	Person   Person
	Bank     Bank
}

type Person struct {
	name    string
	Address string
}

type Bank struct {
	name  string
	Saldo int
}

type interfaceName interface {
	getInformasiName() string
}

func (person Person) getInformasiName() string {
	return person.name
}
func (bank Bank) getInformasiName() string {
	return bank.name
}
func PublicInformasiName(intfName interfaceName) string {
	return intfName.getInformasiName()
}
func CreateUser(user string, name string, email string, nameBank string, saldo int) User {
	return User{user, Person{name, email}, Bank{nameBank, saldo}}
}

type History struct {
	Username  string
	Transaksi string
	Nominal   int
}