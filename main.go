package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!")

// blueprint for struct
type account struct {
	url      string
	login    string
	password string
}

// blueprint for composition struct
type accountWithTimeStamps struct {
	createdAt time.Time
	updatedAt time.Time
	account
}

// method struct
func (acc *accountWithTimeStamps) outputInfo() {
	fmt.Println(acc.url, acc.login, acc.password, acc.createdAt, (*acc).updatedAt)
	fmt.Println(*acc)
}

// method struct
func (acc *account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

// comosition struct
func newMyAccountWithTimeStamps(urlString, login, password string) (*accountWithTimeStamps, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &accountWithTimeStamps{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		account: account{
			url:      urlString,
			login:    login,
			password: password,
		},
	}
	if password == "" {
		newAcc.account.generatePassword(4)
	}
	return newAcc, nil
}

// constructor struct
/*func newAccount(urlString, login, password string) (*account, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	anyAcc := &account{
		url:      urlString,
		login:    login,
		password: password,
	}
	if password == "" {
		anyAcc.generatePassword(12)
	}
	return anyAcc, nil
}*/

func main() {
	url := addInfo("введите url")
	login := addInfo("введите логин")
	password := addInfo("введите пароль")
	myAccount, err := newMyAccountWithTimeStamps(url, login, password)
	if err != nil {
		fmt.Println("неверный формат URL или логин")
		return
	}
	myAccount.outputInfo()
}

func addInfo(data string) string {
	fmt.Print(data, ":")
	var info string
	fmt.Scanln(&info)
	return info
}
