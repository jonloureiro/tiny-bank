package usecases

import (
	"github.com/jonloureiro/tiny-bank/extensions/id"
	"github.com/jonloureiro/tiny-bank/extensions/jwt"
)

type TransferBetweenAccountsInput struct {
	AccessToken          *jwt.Token
	OriginAccountID      id.ID
	DestinationAccountID id.ID
	Amount               int
}

type TransferBetweenAccountsOutput struct {
	TransferID id.ID
}

func (uC *TinyBankUseCases) TransferBetweenAccounts(input *TransferBetweenAccountsInput) (*TransferBetweenAccountsOutput, error) {
	return &TransferBetweenAccountsOutput{
		TransferID: id.New(),
	}, nil
}
