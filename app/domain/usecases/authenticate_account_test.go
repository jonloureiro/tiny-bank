package usecases_test

import (
	"testing"

	"github.com/jonloureiro/tiny-bank/app/domain/entities"
	"github.com/jonloureiro/tiny-bank/app/domain/usecases"
	"github.com/jonloureiro/tiny-bank/app/domain/usecases/repositories/mocks"
	"github.com/jonloureiro/tiny-bank/app/domain/vo"
	"github.com/jonloureiro/tiny-bank/extensions/jwt"
)

func TestAuthenticate(t *testing.T) {
	const (
		validName   = "Test"
		validSecret = "123456"
		privateKey  = "s3cr3t"
	)

	t.Run("", func(t *testing.T) {
		uC := NewTinyBankUsecases()
		cpf, _ := vo.NewCPF(mocks.ValidCPF)
		account, _ := entities.NewAccount(validName, validSecret, cpf, uC.InitialAmount)
		uC.AccountsRepo.Create(account)
		input := usecases.AuthenticateAccountInput{
			CPF:    account.CPF.Value(),
			Secret: account.Secret,
		}
		output, err := uC.AuthenticateAccount(&input)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if output.AccessToken.AccountId != account.ID {
			t.Errorf("expected account id %v, got %v", account.ID, output.AccessToken.AccountId)
		}
		_, err = jwt.Parse(output.AccessToken.Token, privateKey)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})
}
