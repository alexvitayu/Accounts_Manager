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
	//myVault := account.NewVault((cloud.NewCloudDb("https://yandex.ru")))

Menu:
	for {
		userChoise := selectMenu()
		switch userChoise {
		case 1:
			createAccount(myVault)
		case 2:
			findAccount(myVault)
		case 3:
			deleteAccount(myVault)
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

func createAccount(myVault *account.VaultWithDb) {
	url := addInfo("введите url")
	login := addInfo("введите логин")
	password := addInfo("введите пароль")
	myAccount, err := account.NewMyAccount(url, login, password)
	if err != nil {
		output.OutputErrors(err)
		return
	}
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

func findAccount(myVault *account.VaultWithDb) {
	url := addInfo("введите url для поиска")
	accounts := myVault.FindAccountByURL(url)
	if len(accounts) == 0 {
		output.OutputErrors("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.OutputInfo()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := addInfo("введите url аккаунта, который хотите удалить")
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("аккаунт удалён")
	} else {
		color.Red("аккаунт не найден")
	}
}
