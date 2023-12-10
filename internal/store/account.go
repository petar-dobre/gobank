package store

import (
	"database/sql"
	"fmt"

	"github.com/petar-dobre/gobank/internal/models"
)

func (s *PostgresStore) CreateAccount(acc *models.Account) error {
	fmt.Println("Im in!")
	query := `insert into account
  (first_name, last_name, email, password, number, balance, created_at)
  values ($1, $2, $3, $4, $5, $6, $7)`

	fmt.Println(acc)

	resp, err := s.db.Query(query,
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

	fmt.Printf("%v\n", resp)
	return nil
}

func (s *PostgresStore) UpdateAccount(*models.Account) error {
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	_, err := s.db.Query("delete from account where id = $1", id)
	return err
}

func (s *PostgresStore) GetAccounts() ([]*models.AccountDTO, error) {
	rows, err := s.db.Query("select * from account")
	if err != nil {
		return nil, err
	}

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

func (s *PostgresStore) GetAccountByID(id int) (*models.AccountDTO, error) {
	rows, err := s.db.Query(`select 
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

func scanIntoAccount(rows *sql.Rows) (*models.AccountDTO, error) {
	account := new(models.AccountDTO)
	err := rows.Scan(
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