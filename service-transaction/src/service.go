package src

import (
	"time"
)

type ServiceList interface {
	Transaction(data []byte) (int, string, error)
	Reversal(data []byte) (int, string, error)
	CheckPassword(id int, pass string) (int, error)
	CheckUserType(id int) (int, error)
	CheckSufficientBalance(id int, value float64) (int, float64, error)
	CheckApiResponse() (int, error)
	ExecuteTransaction(value, nextBalance float64, payer, payee, types int, transferTime, endDateReversal time.Time) (int, int, error)
	SendNotification(transferID int) (int, error)
	CheckTransferEfetue(id int) (int, error)
	ExecuteReversal(transactionID, payer, payee int) (int, error)
}
