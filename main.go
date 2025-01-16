package main

import (
	"demo/password-1/account"
	"fmt"
)

func main() {
Menu:
	for {
		userChoise := selectMenu()
		switch userChoise {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			deleteAccount()
		default:
			fmt.Println("вы действительно хотите выйти? y/n")
			var ch string
			fmt.Scan(&ch)
			if ch == "y" {
				break Menu
			}
		}
	}

}

func createAccount() {
	url := addInfo("введите url")
	login := addInfo("введите логин")
	password := addInfo("введите пароль")
	myAccount, err := account.NewMyAccount(url, login, password)
	if err != nil {
		fmt.Println("неверный формат URL или логин")
		return
	}
	myVault := account.NewVault()
	myVault.AddAccount(*myAccount)
}

func addInfo(data string) string {
	fmt.Print(data, ":")
	var info string
	fmt.Scanln(&info)
	return info
}

func selectMenu() int {
	fmt.Println("Меню работы с аккаунтом")
	fmt.Println(" 1. Создать аккаунт;")
	fmt.Println(" 2. Найти аккаунт;")
	fmt.Println(" 3. Удалить аккаунт;")
	fmt.Println(" 4. Выход")
	var choise int
	fmt.Scan(&choise)
	return choise
}

func findAccount() {

}

func deleteAccount() {

}
