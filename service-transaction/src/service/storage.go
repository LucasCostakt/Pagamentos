package service

import (
	"service-transaction/src"
	"time"
)

type Service struct {
	Repository src.RepositoryList
}

type Transaction struct {
	Value           float64   `json:"value"`
	Payer           int       `json:"payer"`
	Payee           int       `json:"payee"`
	Password        string    `json:"password"`
	TransferTime    time.Time `json:"transferTime"`
	EndDateReversal time.Time `json:"end_date_reversal"`
}
type ReversalRequest struct {
	UserID     int    `json:"user_id"`
	TransferID int    `json:"transfer_id"`
	Password   string `json:"password"`
}

type Reversal struct {
	Value float64 `json:"value"`
	Payer int     `json:"payer"`
	Payee int     `json:"payee"`
}

type Notification struct {
	Value         float64 `json:"value"`
	Payer_id      int     `json:"payer_id"`
	Payer_name    string  `json:"payer_name"`
	Payee_id      int     `json:"payee_id"`
	Payee_name    string  `json:"payee_name"`
	Transfer_time string  `json:"transfer_time"`
	Email_payer   string  `json:"email_payer"`
	Email_payee   string  `json:"email_payee"`
}
