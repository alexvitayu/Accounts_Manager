package account

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var myLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!")

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

type Account struct {
	Url      string `json:"url"`
	Login    string `json:"ThisAccountLogin"`
	password string
}

type AccountWithTimeStamps struct {
	Account
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time
}

func NewMyAccountWithTimeStamps(urlString, login, password string) (*AccountWithTimeStamps, error) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	if len(login) == 0 {
		return nil, errors.New("INVALID_LOGIN")
	}
	myAccount := &AccountWithTimeStamps{
		Account: Account{
			Url:      urlString,
			Login:    login,
			password: password,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if len(password) == 0 {
		myAccount.generatePassword(8)
	}
	return myAccount, nil
}

/*func NewMyAccount(urlString, login, password string) (*Account, error) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	if len(login) == 0 {
		return nil, errors.New("INVALID_LOGIN")
	}
	myAccount := &Account{
		Url:      urlString,
		Login:    login,
		password: password,
	}
	if len(password) == 0 {
		myAccount.generatePassword(8)
	}
	return myAccount, nil
}*/

func (acc *AccountWithTimeStamps) OutputAccount() {
	color.Yellow(acc.Url)
	color.Yellow(acc.Login)
	color.Yellow(acc.password)
	fmt.Println(acc.CreatedAt, acc.UpdatedAt)
}

func (acc *Account) generatePassword(n int) {
	pass := make([]rune, n)
	for index := range pass {
		pass[index] = myLetters[rand.Intn(len(myLetters))]
	}
	acc.password = string(pass)
}
