package store

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/petar-dobre/gobank/internal/models"
)

type Storage interface {
	CreateAccount(*models.Account) error
	DeleteAccount(int) error
	UpdateAccount(*models.Account) error
	GetAccounts() ([]*models.AccountDTO, error)
	GetAccountByID(int) (*models.AccountDTO, error)
	GetHashedPassword(string) (string, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, err
}

func (s *PostgresStore) Init() error {
	return s.CreateAccountTable()
}

func (s *PostgresStore) CreateAccountTable() error {
	query := `create table if not exists account (
    id serial primary key,
    first_name varchar(50),
    last_name varchar(50),
		email varchar(50),
		password varchar,
    number serial,
    balance serial,
    created_at timestamp
  )`

	_, err := s.db.Exec(query)

	return err
}
