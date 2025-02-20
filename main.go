package main

import (
	"fmt"
	"revision/part-1/account"
	"revision/part-1/files"
	"revision/part-1/output"
	"revision/part-1/reverse"
	"strings"

	"github.com/fatih/color"
)

// Reverse array | mutation of the origin object by using pointers | generic for different types
// Description struct and create instance
// transfer struct
// apply pointers
// rune struct
// generate password
// method struct
// mutation struct
// constructor struct
// validation of the data
// transfer of the generation
// composition struct
// split the package
// add the package
// import/export
// add the external package
// files package & go mog tidy
// Write file & stack frame & defer
// read file
// json & struct tags & save json
// select menu exercise
// slice struct | vault
// read json
// searching the password exercise
// deleting the password exercise
// interface: change files
// dependencies injection
// the second provider
// creation interface
// embedded interface
// any type
// type switch
// obtain type
// generic
// hack in generic restrictions
// generic struct
// input generic exercise
// Advanced functions & map instead of switch-case
// transfer of function
// amonymus functions
// searching by login exercise
// dynamic number of arguments
// closure

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

func main() {

	// Reverse array | mutation of the origin object by using pointers | generic for different types

	r := []int{1, 2, 3, 4, 5, 6}
	//r := []string{"a", "b", "c", "d", "e", "f"}
	//r := []float64{0.5, 0.6, 0.7, 0.8, 0.9, 1.0}
	reverse.OutputReverse(&r)
	rev := reverse.ReverseArray(&r)
	reverse.OutputReverse(rev)

	myVault := account.NewMyVault(files.NewJsondb("data.json"))
	//myVault := account.NewMyVault(cloud.NewcloudDb("https://yandex.ru"))

Menu:
	for {
		choise := promptData(
			"_Меню работы с аккаунтами_",
			"1. Создать аккаунт",
			"2. Найти аккаунт по Url",
			"3. Найти аккаунт по Login",
			"4. Удалить аккаунт",
			"5. Выход",
			"Выберите вариант",
		)
		selectMenu := menu[choise]
		if selectMenu == nil {
			break Menu
		}
		selectMenu(myVault)
		/*switch choise {
		case "1":
			createAccount(myVault)
		case "2":
			findAccount(myVault)
		case "3":
			deleteAccount(myVault)
		default:
			fmt.Println("Вы хотите выйти?, y/n")
			var ch string
			fmt.Scan(&ch)
			if ch == "y" {
				break Menu
			}
		}*/
	}

}

func createAccount(vault *account.VaultWithDb) {
	// Work with Accounts&Vault
	url := promptData("Введите url")
	login := promptData("Введите login")
	password := promptData([]string{"Введите password"})
	myAccount, err := account.NewMyAccountWithTimeStamps(url, login, password)
	if err != nil {
		output.OutputErrorHack[error](err)
		return
	}
	err = vault.AddAccount(*myAccount)
	if err != nil {
		output.OutputErrorHack(err)
		return
	}
	vault.AddAccount(*myAccount)
}

func promptData(data ...any) string {
	for index, value := range data {
		if index == len(data)-1 {
			fmt.Printf("%v: ", value)
		} else {
			fmt.Println(value)
		}
	}
	var input string
	fmt.Scanln(&input)
	return input
}

func findAccountByUrl(vault *account.VaultWithDb) {
	url := promptData("введите url для поиска аккаунта")
	accounts := vault.FindAccounts(url, func(str string, acc account.AccountWithTimeStamps) bool {
		return strings.Contains(acc.Url, str)
	})
	outPut(accounts)
}
func findAccountByLogin(vault *account.VaultWithDb) {
	login := promptData("введите login для поиска аккаунта")
	accounts := vault.FindAccounts(login, func(str string, acc account.AccountWithTimeStamps) bool {
		return strings.Contains(acc.Login, str)
	})
	outPut(accounts)
}

func outPut(accounts *[]account.AccountWithTimeStamps) {
	if len(*accounts) == 0 {
		output.OutputErrorHack("No found accounts")
	}
	for _, account := range *accounts {
		account.OutputAccount()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData("введите url чтобы удалить аккаунт")
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("Аккаунт успешно удалён")
	} else {
		color.Red("Аккаунты не найдены")
	}
}
