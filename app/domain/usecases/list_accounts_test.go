package usecases_test

import (
	"testing"

	"github.com/jonloureiro/tiny-bank/app/domain/entities"
	"github.com/jonloureiro/tiny-bank/app/domain/usecases"
	"github.com/jonloureiro/tiny-bank/app/domain/usecases/repositories/mocks"
	"github.com/jonloureiro/tiny-bank/app/domain/vo"
)

func TestListAccounts(t *testing.T) {
	var (
		validName     = "Test"
		validSecret   = "123456"
		validCPF, _   = vo.NewCPF("69029890100")
		initialAmount = 100
	)
	t.Run("list accounts", func(t *testing.T) {
		uC := NewTinyBankUsecases()
		account, _ := entities.NewAccount(validName, validSecret, validCPF, initialAmount)
		_ = uC.AccountsRepo.Create(account)
		output, err := uC.ListAccounts(usecases.ListAccountsInput{})
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if len(output.Accounts) != 1 {
			t.Errorf("expected 1 account, got %d", len(output.Accounts))
		}
	})

	t.Run("validate empty slice", func(t *testing.T) {
		want := make([]*entities.Account, 0)
		uC := NewTinyBankUsecases()
		output, err := uC.ListAccounts(usecases.ListAccountsInput{})
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if len(output.Accounts) != len(want) {
			t.Errorf("expected %d accounts, got %d", len(want), len(output.Accounts))
		}
	})

	t.Run("validate error", func(t *testing.T) {
		uC := NewTinyBankUsecases()
		accountsRepo := mocks.NewAccountsRepositoryMock()
		accountsRepo.UnknownError = true
		uC.AccountsRepo = accountsRepo
		_, err := uC.ListAccounts(usecases.ListAccountsInput{})
		if err != usecases.ErrDatabaseUnknownError {
			t.Errorf("expected error %v, got %v", usecases.ErrDatabaseUnknownError, err)
		}
	})
}
