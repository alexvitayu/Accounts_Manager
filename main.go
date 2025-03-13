package main

import (
	"demo/password-1/account"
	"demo/password-1/encryptor"
	"demo/password-1/files"
	"demo/password-1/output"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

var menuVariants = []string{
	" 1. Создать аккаунт;",
	" 2. Найти аккаунт по url;",
	" 3. Найти аккаунт по login;",
	" 4. Удалить аккаунт;",
	" 5. Выход",
	"выберите вариант",
}

func main() {
	fmt.Println("_Меню работы с аккаунтом_")
	err := godotenv.Load()
	if err != nil {
		output.OutputErrors("не удалось найти env-файл")
	}
	myVault := account.NewVault(files.NewJsonDb("data.vault"), *encryptor.NewEncryptor())
	//myVault := account.NewVault((cloud.NewCloudDb("https://yandex.ru")))

Menu:
	for {
		variant := promptData(menuVariants...)
		menuFunc := menu[variant]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(myVault)
		/*switch variant {
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
		}*/

	}
}

func createAccount(myVault *account.VaultWithDb) {
	url := promptData("введите url")
	login := promptData("введите логин")
	password := promptData("введите пароль")
	myAccount, err := account.NewMyAccount(url, login, password)
	if err != nil {
		output.OutputErrors(err)
		return
	}
	myVault.AddAccount(*myAccount)
}

// функция принимает slice любого типа
// выводит строкой каждый элемент, а последний - Printf добавляя :
func promptData(prompt ...string) string {
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

func findAccountByUrl(myVault *account.VaultWithDb) {
	url := promptData("введите url для поиска")
	accounts := myVault.FindAccounts(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})
	outPut(&accounts)
}

func findAccountByLogin(myVault *account.VaultWithDb) {
	login := promptData("введите login для поиска")
	accounts := myVault.FindAccounts(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})
	outPut(&accounts)
}

func outPut(acc *[]account.Account) {
	if len(*acc) == 0 {
		output.OutputErrors("Аккаунтов не найдено")
	}
	for _, account := range *acc {
		account.OutputInfo()
	}

}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData("введите url аккаунта, который хотите удалить")
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("аккаунт удалён")
	} else {
		color.Red("аккаунт не найден")
	}
}
