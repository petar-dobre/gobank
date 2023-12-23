package models

type PostgresStorer interface {
	AccountStorer
}

type AuthStorer interface {
	GetHashedPassword(string) (string, error)
}

type AccountStorer interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*UpdateAccountReq, int) error
	GetAccounts() ([]*AccountDTO, error)
	GetAccountByID(int) (*AccountDTO, error)
}
