package main

import (
	"fmt"
	"math/rand/v2"
)

type account struct {
	url      string
	login    string
	password string
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!")

func main() {

	url := addInfo("введите url")
	login := addInfo("введите логин")
	fmt.Println("Вы хотите сгенерировать пароль? y/n:")
	var userChoise string
	var password string
	fmt.Scan(&userChoise)
	switch userChoise {
	case "y":
		password = generatePassword(8)
	case "n":
		password = addInfo("введите пароль")
	}

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

func generatePassword(n int) string {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	return string(res)
}
