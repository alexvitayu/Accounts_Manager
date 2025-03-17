package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var myLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!")

type Account struct {
	Url      string `json:"url"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AccountWithTimeStamps struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Account
}

func (acc *AccountWithTimeStamps) OutputInfo() {
	color.Green(acc.Url)
	color.Green(acc.Login)
	color.Green(acc.Password)
	fmt.Println(acc.CreatedAt, acc.UpdatedAt)

}

func (acc *AccountWithTimeStamps) generatePassword(n int) {
	pas := make([]rune, n)
	for i := range pas {
		pas[i] = myLetters[rand.IntN(len(myLetters))]
	}
	acc.Account.Password = string(pas)
}

func NewMyAccountWithTimeStamps(urlString, login, password string) (*AccountWithTimeStamps, error) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	myAccount := &AccountWithTimeStamps{
		Account: Account{
			Url:      urlString,
			Login:    login,
			Password: password,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if password == "" {
		myAccount.generatePassword(8)
	}
	return myAccount, nil
}

/*func NewMyAccount(urlString, login, password string) (*account, error) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	myAccount := &account{
		url:      urlString,
		login:    login,
		password: password,
	}
	if password == "" {
		myAccount.generatePassword(8)
	}
	return myAccount, nil
}*/
