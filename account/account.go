package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!")

// blueprint for struct
type Account struct {
	url      string
	login    string
	password string
}

// blueprint for composition struct
type AccountWithTimeStamps struct {
	createdAt time.Time
	updatedAt time.Time
	Account
}

// method struct
func (acc *AccountWithTimeStamps) OutputInfo() {
	fmt.Println(acc.url, acc.login, acc.password, acc.createdAt, (*acc).updatedAt)
	fmt.Println(*acc)
}

// method struct
func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

// comosition struct
func NewMyAccountWithTimeStamps(urlString, login, password string) (*AccountWithTimeStamps, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &AccountWithTimeStamps{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		Account: Account{
			url:      urlString,
			login:    login,
			password: password,
		},
	}
	if password == "" {
		newAcc.Account.generatePassword(4)
	}
	return newAcc, nil
}
