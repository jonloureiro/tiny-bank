package usecases_test

import (
	"testing"

	"github.com/jonloureiro/tiny-bank/app/domain/usecases"
	"github.com/jonloureiro/tiny-bank/extensions/jwt"
)

func TestTransferBetweenAccounts(t *testing.T) {
	const (
		privateKey  = "s3cr3t"
		validSecret = "123456"
		validCPF1   = "02561075133"
		validCPF2   = "61003251102"
	)

	t.Run("create transfer", func(t *testing.T) {
		uC := NewTinyBankUsecases()
		originAccountOutput, _ := uC.CreateAccount(usecases.CreateAccountInput{"origin", validSecret, validCPF1})
		destinationAccountOutput, _ := uC.CreateAccount(usecases.CreateAccountInput{"destination", validSecret, validCPF2})
		token, _ := jwt.New(originAccountOutput.AccountID, privateKey)
		output, err := uC.TransferBetweenAccounts(&usecases.TransferBetweenAccountsInput{
			AccessToken:          token,
			OriginAccountID:      originAccountOutput.AccountID,
			DestinationAccountID: destinationAccountOutput.AccountID,
			Amount:               100,
		})
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if output.TransferID == "" {
			t.Errorf("expected transfer id, got empty string")
		}
	})
}
