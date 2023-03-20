package usecases

import "github.com/jonloureiro/tiny-bank/app/domain/entities"

type ListAccountsInput struct {
}

type ListAccountsOutput struct {
	Accounts []*entities.Account
}

func (uC *TinyBankUseCases) ListAccounts(input ListAccountsInput) (*ListAccountsOutput, error) {
	accounts, err := uC.AccountsRepo.List()
	if err != nil {
		return nil, ErrDatabaseUnknownError
	}
	return &ListAccountsOutput{Accounts: accounts}, nil
}
