package main

import (
	"app/controllers"
	"app/models"
	"app/views"
	"fmt"
)
var database []models.User

func main() {
    for{
        views.MainMenuView()
        var input string
        fmt.Scanln(&input)
        validasi := controllers.ValidasiMenu(input, 3)
        if validasi != nil {
            fmt.Println(validasi)
            continue
        }
        if input == "0" {
            break
        }else if input == "1" {
            controllers.ClearScreen()
		    fmt.Println("Halaman Login")
            controllers.Login(&database)
        }else if input == "2" {
            controllers.ClearScreen()
		    fmt.Println("Halaman Registrasi")
            controllers.RegisterUser(&database)
        }else if input == "3" {
            controllers.ClearScreen()
            views.PrintAllUser(database)
        }
    }
}

func init() {
    controllers.ClearScreen()
	fmt.Println("Selamat Datang di App Mobile Banking\nSilahkan pilih menu :")
    user1 := models.CreateUser("user1", "Lumoshive", "kedoya", "BCA", 100000)
    user2 := models.CreateUser("user2", "Academy", "kedoya", "BRI", 100000)
    database = append(database, user1)
    database = append(database, user2)
}