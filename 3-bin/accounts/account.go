package accounts

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ-*!")

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewAccount(login, password, urlNew string) (*Account, error) {
	_, err := url.ParseRequestURI(urlNew)
	if err != nil {
		return nil, errors.New("invalid url")
	}
	if login == "" {
		return nil, errors.New("invalid login")
	}
	newAccount := &Account{
		Login:     login,
		Password:  password,
		Url:       urlNew,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if password == "" {
		newAccount.generatePassword(12)
	}
	return newAccount, nil
}

func (a *Account) generatePassword(n int) {
	data := make([]rune, n)
	for i := range data {
		data[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	a.Password = string(data)
}
