package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
)

type account struct {
	url      string
	login    string
	password string
}

func (acc *account) outputInfo() {
	fmt.Println((acc).url, (acc).login, (acc).password)
}

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

func newAccount(login, password, urlString string) (*account, error) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	return &account{
		url:      urlString,
		login:    login,
		password: password,
	}, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!")

func main() {

	url := addInfo("введите url")
	login := addInfo("введите логин")
	password := addInfo("введите пароль")
	account1, err := newAccount(login, password, url)
	if err != nil {
		fmt.Println("неверный формат URL")
		return
	}
	account1.generatePassword(8)

	account1.outputInfo()
}

func addInfo(data string) string {
	fmt.Print(data, ": ")
	var info string
	fmt.Scan(&info)
	return info
}
