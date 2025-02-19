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
// Advanced functions - map intead of switch-case
// transfer of functions
// anonimous functions & searching by login & delete the duplicated code
// dynamic number of arguments

var variantMenu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccount,
	"3": deleteAccount,
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
		choise := promptData([]string{
			"_Меню работы с аккаунтами_",
			"1. Создать аккаунт",
			"2. Найти аккаунт",
			"3. Удалить аккаунт",
			"4. Выход",
			"Выберите вариант",
		})
		menuSelect := variantMenu[choise]
		if menuSelect == nil {
			break Menu
		}
		menuSelect(myVault)
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
	url := promptData([]string{"Введите url"})
	login := promptData([]string{"Введите login"})
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

func promptData[T any](data []T) string {
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

func findAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"введите url для поиска аккаунта"})
	accounts := vault.FindAccounts(url, finder)
	if len(*accounts) == 0 {
		output.OutputErrorHack("Аккаунтов не найдено")
	}
	for _, account := range *accounts {
		account.OutputAccount()
	}
}

func finder(str string, acc account.AccountWithTimeStamps) bool {
	return strings.Contains(acc.Url, str)
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"введите url чтобы удалить аккаунт"})
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("Аккаунт успешно удалён")
	} else {
		color.Red("Аккаунты не найдены")
	}
}
