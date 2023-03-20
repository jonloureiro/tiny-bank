package entities_test

import (
	"testing"

	"github.com/jonloureiro/tiny-bank/app/domain/entities"
	"github.com/jonloureiro/tiny-bank/app/domain/vo"
	"github.com/jonloureiro/tiny-bank/extensions/jwt"
)

func TestNewAccount(t *testing.T) {
	var (
		validCPF, _   = vo.NewCPF("69029890100")
		initialAmount = 100
	)

	t.Run("create account", func(t *testing.T) {
		account, err := entities.NewAccount("Jon", "123456", validCPF, initialAmount)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if account == nil {
			t.Errorf("expected an account, got %v", account)
		}
	})

	t.Run("validate errors", func(t *testing.T) {
		testCases := map[string]struct {
			name        string
			cpf         *vo.CPF
			secret      string
			expectedErr error
		}{
			"empty name": {
				name:        "",
				cpf:         validCPF,
				secret:      "123456",
				expectedErr: entities.ErrEmptyName,
			},
			"empty secret": {
				name:        "Jon",
				cpf:         validCPF,
				secret:      "",
				expectedErr: entities.ErrEmptySecret,
			},
			"empty cpf": {
				name:        "Jon",
				cpf:         nil,
				secret:      "123456",
				expectedErr: entities.ErrEmptyCPF,
			},
		}
		for desc, tC := range testCases {
			t.Run(desc, func(t *testing.T) {
				account, err := entities.NewAccount(tC.name, tC.secret, tC.cpf, initialAmount)
				if err != tC.expectedErr {
					t.Errorf("expected error %v, got %v", tC.expectedErr, err)
				}
				if account != nil {
					t.Errorf("expected no account, got %v", account)
				}
			})
		}
	})

}

func TestAuthenticateAccount(t *testing.T) {
	var (
		validCPF, _   = vo.NewCPF("69029890100")
		initialAmount = 100
	)

	privateKey := "s3cr3t"
	secret := "123456"
	account, err := entities.NewAccount("Jon", secret, validCPF, initialAmount)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	got, _ := account.Authenticate(secret, privateKey)
	want, _ := jwt.Parse(got.Token, privateKey)
	if got.Token != want.Token {
		t.Errorf("expected token %v, got %v", want.Token, got.Token)
	}
}
