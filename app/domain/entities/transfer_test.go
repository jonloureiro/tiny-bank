package entities_test

import (
	"testing"

	"github.com/jonloureiro/tiny-bank/app/domain/entities"
	"github.com/jonloureiro/tiny-bank/app/domain/vo"
)

func TestNewTransfer(t *testing.T) {
	const (
		validCPF1     = "02561075133"
		validCPF2     = "61003251102"
		validSecret   = "123456"
		initialAmount = 1000
	)

	t.Run("create transfer", func(t *testing.T) {
		cpf1, _ := vo.NewCPF(validCPF1)
		cpf2, _ := vo.NewCPF(validCPF2)
		originAccount, _ := entities.NewAccount("origin", validSecret, cpf1, initialAmount)
		destinationAccount, _ := entities.NewAccount("destination", validSecret, cpf2, initialAmount)
		amount := 100
		transfer, err := entities.NewTransfer(originAccount, destinationAccount, amount)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if transfer.ID == "" {
			t.Errorf("expected transfer id, got empty string")
		}
		if transfer.OriginAccount.ID == "" {
			t.Errorf("expected origin account id, got empty string")
		}
		if transfer.DestinationAccount.ID == "" {
			t.Errorf("expected destination account id, got empty string")
		}
		if transfer.Amount != amount {
			t.Errorf("expected amount 0, got %d", transfer.Amount)
		}
	})

}
