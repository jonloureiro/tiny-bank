package repositories

import (
	"github.com/jonloureiro/tiny-bank/app/domain/entities"
	"github.com/jonloureiro/tiny-bank/extensions/id"
)

type TransfersRepository interface {
	List(accountID id.ID) ([]*entities.Transfer, error)
	Create(transfer *entities.Transfer) error
}
