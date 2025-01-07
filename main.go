package main

import (
	"demo/password-1/account"
	"demo/password-1/files"
	"fmt"
)

func main() {
	createAccount()
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
	file, err := myAccount.ToBytes()
	if err != nil {
		fmt.Println("Не удалось преобразовать в JSON")
		return
	}
	files.WriteFile(file, "data.json")
}

func addInfo(data string) string {
	fmt.Print(data, ":")
	var info string
	fmt.Scanln(&info)
	return info
}
