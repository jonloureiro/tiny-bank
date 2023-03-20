package entities

import (
	"time"

	"github.com/jonloureiro/tiny-bank/extensions/id"
)

type Transfer struct {
	ID                 id.ID
	OriginAccount      *Account
	DestinationAccount *Account
	Amount             int
	CreatedAt          time.Time
}

func NewTransfer(originAccount, destinationAccount *Account, amount int) (*Transfer, error) {
	transfer := Transfer{
		ID:                 id.New(),
		OriginAccount:      originAccount,
		DestinationAccount: destinationAccount,
		Amount:             amount,
		CreatedAt:          time.Now(),
	}
	return &transfer, nil
}
