package main

import (
	"demo/password-1/account"
	"demo/password-1/files"
	"demo/password-1/output"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	myVault := account.NewVault(files.NewJsonDb("data.json"))

Menu:
	for {
		choise := getMenu()
		switch choise {
		case 1:
			createAccount(myVault)
		case 2:
			findAccount(myVault)
		case 3:
			deleteAccount(myVault)
		default:
			fmt.Println("Хотите выйти? y/n")
			var ch string
			fmt.Scan(&ch)
			if ch == "y" {
				break Menu
			}
		}
	}
}

func createAccount(vault *account.VaultWithDb) {
	url := addInfo("введите url")
	login := addInfo("введите логин")
	password := addInfo("введите пароль")
	myAccount, err := account.NewMyAccount(url, login, password)
	if err != nil {
		output.OutputErrors(err)
		return
	}
	vault.AddAccount(*myAccount)
}

func addInfo(data string) string {
	fmt.Print(data, ":")
	var info string
	fmt.Scanln(&info)
	return info
}

func getMenu() int {
	fmt.Println("_Менеджер аккаунтов_")
	fmt.Println("Выберите вариант:")
	fmt.Println("1 Создать аккаунт:")
	fmt.Println("2 Найти аккаунт:")
	fmt.Println("3 Удалить аккаунт:")
	fmt.Println("4 Выход:")
	var choise int
	fmt.Scan(&choise)
	return choise
}

func findAccount(vault *account.VaultWithDb) {
	url := addInfo("Введите url для поиска аккаунта:")
	accounts := vault.FindAccountByUrl(url)
	for _, account := range accounts {
		account.OutputInfo()
	}
	if len(accounts) == 0 {
		output.OutputErrors("Аккаунтов не найдено")
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := addInfo("Введите url для поиска аккаунта:")
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("Аккаунт удалён")
	} else {
		color.Red("Аккаунты не найдены")
	}

}
