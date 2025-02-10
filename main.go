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
		choise := promptData[string]([]string{
			"_Менеджер аккаунтов_",
			"1 Создать аккаунт",
			"2 Найти аккаунт",
			"3 Удалить аккаунт",
			"4 Выход",
			"Выберите вариант",
		})
		switch choise {
		case "1":
			createAccount(myVault)
		case "2":
			findAccount(myVault)
		case "3":
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
	url := promptData[string]([]string{"введите url"})
	login := promptData[string]([]string{"введите login"})
	password := promptData[string]([]string{"введите password"})
	myAccount, err := account.NewMyAccount(url, login, password)
	if err != nil {
		output.OutputErrors(err)
		return
	}
	vault.AddAccount(*myAccount)
}

func promptData[T any](data []string) string {
	for i, str := range data {
		if i == len(data)-1 {
			fmt.Printf("%v: ", str)
		} else {
			fmt.Println(str)
		}
	}
	var choise string
	fmt.Scanln(&choise)
	return choise
}

func findAccount(vault *account.VaultWithDb) {
	url := promptData[string]([]string{"введите url для поиска аккаунта"})
	accounts := vault.FindAccountByUrl(url)
	for _, account := range accounts {
		account.OutputInfo()
	}
	if len(accounts) == 0 {
		output.OutputErrors("Аккаунтов не найдено")
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData[string]([]string{"введите url для поиска аккаунта"})
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("Аккаунт удалён")
	} else {
		color.Red("Аккаунты не найдены")
	}

}
