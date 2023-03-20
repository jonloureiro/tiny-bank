package usecases_test

import (
	"github.com/jonloureiro/tiny-bank/app/domain/usecases"
	"github.com/jonloureiro/tiny-bank/app/domain/usecases/repositories/mocks"
)

const PrivateKey = "s3cr3t"

func NewTinyBankUsecases() *usecases.TinyBankUseCases {
	return &usecases.TinyBankUseCases{
		PrivateKey: PrivateKey,

		AccountsRepo: mocks.NewAccountsRepositoryMock(),
		TransferRepo: mocks.NewTransfersRepositoryMock(),
	}
}
