package main

import (
	"demo/password-1/account"
	"fmt"
)

func main() {
	url := addInfo("введите url")
	login := addInfo("введите логин")
	password := addInfo("введите пароль")
	myAccount, err := account.NewMyAccountWithTimeStamps(url, login, password)
	if err != nil {
		fmt.Println("неверный формат URL или логин")
		return
	}
	myAccount.OutputInfo()
}

func addInfo(data string) string {
	fmt.Print(data, ":")
	var info string
	fmt.Scanln(&info)
	return info
}
