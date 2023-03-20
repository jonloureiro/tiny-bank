package entities

import (
	"errors"
	"time"

	"github.com/jonloureiro/tiny-bank/app/domain/vo"
	"github.com/jonloureiro/tiny-bank/extensions/id"
	"github.com/jonloureiro/tiny-bank/extensions/jwt"
)

var (
	ErrEmptyName          = errors.New("empty name")
	ErrEmptySecret        = errors.New("empty secret")
	ErrEmptyCPF           = errors.New("empty cpf")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type Account struct {
	ID            id.ID
	Name          string
	CPF           *vo.CPF
	Secret        string
	initialAmount int
	CreatedAt     time.Time
}

func NewAccount(name, secret string, cpf *vo.CPF, initialAmount int) (*Account, error) {
	if cpf == nil {
		return nil, ErrEmptyCPF
	}
	account := Account{
		ID:            id.New(),
		Name:          name,
		CPF:           cpf,
		Secret:        secret,
		initialAmount: initialAmount,
		CreatedAt:     time.Now(),
	}
	err := account.validate()
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (a *Account) validate() error {
	if a.Name == "" {
		return ErrEmptyName
	}
	if a.Secret == "" {
		return ErrEmptySecret
	}
	return nil
}

func (a *Account) Authenticate(secret, privateKey string) (*jwt.Token, error) {
	if a.Secret != secret {
		return nil, ErrInvalidCredentials
	}
	token, _ := jwt.New(a.ID, privateKey)
	return token, nil
}
