package mocks

import (
	"github.com/jonloureiro/tiny-bank/app/domain/entities"
	"github.com/jonloureiro/tiny-bank/extensions/id"
)

type TransfersRepositoryMock struct {
	StorageByAccountId map[id.ID][]*entities.Transfer
	UnknownError       bool
}

func NewTransfersRepositoryMock() *TransfersRepositoryMock {
	return &TransfersRepositoryMock{
		StorageByAccountId: map[id.ID][]*entities.Transfer{},
	}
}

func (r *TransfersRepositoryMock) List(accountID id.ID) ([]*entities.Transfer, error) {
	if r.StorageByAccountId[accountID] == nil {
		r.StorageByAccountId[accountID] = []*entities.Transfer{}
	}
	transfers, ok := r.StorageByAccountId[accountID]
	if !ok {
		return []*entities.Transfer{}, nil
	}
	return transfers, nil

}

func (r *TransfersRepositoryMock) Create(transfer *entities.Transfer) error {
	if r.StorageByAccountId[transfer.OriginAccount.ID] == nil {
		r.StorageByAccountId[transfer.OriginAccount.ID] = []*entities.Transfer{}
	}
	r.StorageByAccountId[transfer.OriginAccount.ID] = append(r.StorageByAccountId[transfer.OriginAccount.ID], transfer)
	r.StorageByAccountId[transfer.DestinationAccount.ID] = append(r.StorageByAccountId[transfer.DestinationAccount.ID], transfer)
	return nil
}
