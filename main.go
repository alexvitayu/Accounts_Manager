package main

import "fmt"

type account struct {
	url      string
	login    string
	password string
}

func main() {

	url := addInfo("введите url")
	login := addInfo("введите логин")
	password := addInfo("введите пароль")

	account1 := account{
		login:    login,
		url:      url,
		password: password,
	}

	outputInfo(&account1)
}

func addInfo(info string) string {
	fmt.Print(info, ": ")
	var num string
	fmt.Scan(&num)
	return num
}

func outputInfo(acc *account) {
	fmt.Println(acc)
	fmt.Println((*acc).url, acc.login, acc.password)
}
