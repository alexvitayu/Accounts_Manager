package main

import (
	"fmt"
	"revision/part-1/account"
	"revision/part-1/encryptor"
	"revision/part-1/files"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
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
// get env and reading of .env-file
// ENCRYPTER struct and methods Encrypt and Decrypt
// Application of encrypter

var variants = []string{
	"_Меню работы с аккаунтом_",
	"1 создать аккаунт,",
	"2 найти аккаунт по url,",
	"3 найти аккаунт по логин,",
	"4 удалить аккаунт,",
	"5 выход",
	"Выберите вариант",
}

func main() {
	myMenu := map[string]func(*account.VaultWithDb){
		"1": createAccount,
		"2": findAccountByUrl,
		"3": findAccountByLogin,
		"4": deleteAccount,
	}

	godotenv.Load()

	myVault := account.NewMyVault(files.NewJsonDb("data.vault"), encryptor.NewEncryptor())
	//myVault := account.NewMyVault(cloud.NewCloudDb("https://yandex.ru"))

Menu:
	for {
		color.Yellow("Всего %d аккаунтов", len(myVault.Accounts))
		variant := promptData(variants...)

		menu := myMenu[variant]
		if menu == nil {
			fmt.Println("Хотите выйти? y/n")
			var ch string
			fmt.Scan(&ch)
			if ch == "y" {
				break Menu
			}
		}
		menu(myVault)
	}
}

/*
   Menu:

   	switch variant {
   	case "1":
   		createAccount(myVault)
   	case "2":
   		findAccount(myVault)
   	case "3":
   		deleteAccount(myVault)
   	case "4":
   		fmt.Println("Хотите выйти? y/n")
   		var ch string
   		fmt.Scan(&ch)
   		if ch == "y" {
   			break Menu
   		}
   	}
*/

func createAccount(vault *account.VaultWithDb) {
	url := promptData("Введите url")
	login := promptData("Введите login")
	password := promptData("Введите password")

	myAccount, err := account.NewMyAccountWithTimeStamps(url, login, password)
	if err != nil {
		fmt.Println(err)
		return
	}
	vault.AddAccount(myAccount)

}

/*func inputData(inp string) string {
	fmt.Print(inp, ": ")
	var input string
	fmt.Scanln(&input)
	return input
}*/

func promptData(data ...string) string {
	for i, str := range data {
		if i == len(data)-1 {
			fmt.Print(str, ": ")
		} else {
			fmt.Println(str)
		}
	}
	var choise string
	fmt.Scanln(&choise)
	return choise
}

/*func selectMenu() int {
	var ch int
	fmt.Println("_Меню работы с аккаунтом_")
	fmt.Println("1 создать аккаунт,")
	fmt.Println("2 найти аккаунт,")
	fmt.Println("3 удалить аккаунт,")
	fmt.Println("4 выход")
	fmt.Scan(&ch)
	return ch
}*/

func findAccountByUrl(vault *account.VaultWithDb) {
	url := promptData("Введите url")
	accounts := vault.FindAccount(url, func(url string, acc account.AccountWithTimeStamps) bool {
		return strings.Contains(acc.Url, url)
	})
	outPut(accounts)
}

func findAccountByLogin(vault *account.VaultWithDb) {
	login := promptData("Введите login")
	accounts := vault.FindAccount(login, func(login string, acc account.AccountWithTimeStamps) bool {
		return strings.Contains(acc.Login, login)
	})
	outPut(accounts)
}

func outPut(accounts *[]account.AccountWithTimeStamps) {
	if len(*accounts) > 0 {
		color.Green("Найдено %d аккуаунтов", len(*accounts))
	} else {
		color.Red("не найдено аккаунтов")
	}
	for _, acc := range *accounts {
		acc.OutputInfo()
	}

}

/*func checker(url string, acc account.AccountWithTimeStamps) bool {
	return strings.Contains(acc.Url, url)
}*/

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData("Введите url")
	isCompared := vault.DeleteAccountByUrl(url)
	if isCompared {
		color.Green("Аккаунт удалён")
	} else {
		color.Red("не найдено аккаунтов")
	}
}
