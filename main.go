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
		userChoise := promptData([]string{
			"_Меню работы с аккаунтом_",
			" 1. Создать аккаунт;",
			" 2. Найти аккаунт;",
			" 3. Удалить аккаунт;",
			" 4. Выход",
			"выберите вариант",
		})
		switch userChoise {
		case "1":
			createAccount(myVault)
		case "2":
			findAccount(myVault)
		case "3":
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
	url := promptData([]string{"введите url"})
	login := promptData([]string{"введите логин"})
	password := promptData([]string{"введите пароль"})
	myAccount, err := account.NewMyAccount(url, login, password)
	if err != nil {
		output.OutputErrors(err)
		return
	}
	myVault.AddAccount(*myAccount)
}

// функция принимает slice любого типа
// выводит строкой каждый элемент, а последний - Printf добавляя :
func promptData[T any](prompt []T) string {
	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v:", line)
		} else {
			fmt.Println(line)
		}
	}
	var info string
	fmt.Scanln(&info)
	return info
}

func findAccount(myVault *account.VaultWithDb) {
	url := promptData([]string{"введите url для поиска"})
	accounts := myVault.FindAccountByURL(url)
	if len(accounts) == 0 {
		output.OutputErrors("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.OutputInfo()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"введите url аккаунта, который хотите удалить"})
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("аккаунт удалён")
	} else {
		color.Red("аккаунт не найден")
	}
}
