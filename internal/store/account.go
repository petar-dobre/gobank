package store

import (
	"database/sql"
	"fmt"

	"github.com/petar-dobre/gobank/internal/models"
)

type AccountStore struct {
	*PostgresStore
}

func (as *AccountStore) CreateAccount(acc *models.Account) error {
	query := `insert into account
  (first_name, last_name, email, password, number, balance, created_at)
  values ($1, $2, $3, $4, $5, $6, $7)`

	_, err := as.db.Query(query,
		acc.FirstName,
		acc.LastName,
		acc.Email,
		acc.Password,
		acc.Number,
		acc.Balance,
		acc.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (as *AccountStore) UpdateAccount(acc *models.UpdateAccountReq, id int) error {
	query := `update account
	set first_name = $1, last_name = $2
  where id = $3`
	_, err := as.db.Query(query, acc.FirstName, acc.LastName, id)

	if err != nil {
		return err
	}

	return nil
}

func (as *AccountStore) DeleteAccount(id int) error {
	_, err := as.db.Query("delete from account where id = $1", id)
	return err
}

func (as *AccountStore) GetAccounts() ([]*models.AccountDTO, error) {
	rows, err := as.db.Query("select id, first_name, last_name, email, number, balance, created_at  from account")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	
	accounts := []*models.AccountDTO{}
	for rows.Next() {
		account, err := scanIntoAccount(rows)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (as *AccountStore) GetAccountByID(id int) (*models.AccountDTO, error) {
	rows, err := as.db.Query(`select 
	id, first_name, last_name, email, number, balance, created_at 
	from account where id = $1`, id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoAccount(rows)
	}

	return nil, fmt.Errorf("account %d not found", id)
}

func scanIntoAccount(row *sql.Rows) (*models.AccountDTO, error) {
	account := new(models.AccountDTO)
	err := row.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Email,
		&account.Number,
		&account.Balance,
		&account.CreatedAt,
	)

	return account, err
}
