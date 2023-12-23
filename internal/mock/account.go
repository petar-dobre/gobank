package mock

import (
	"fmt"
	"time"

	"github.com/petar-dobre/gobank/internal/models"
)

type MockAccountStore struct {
	MockAccounts []*models.AccountDTO
}

func (ms *MockAccountStore) GetAccounts() ([]*models.AccountDTO, error) {
	return ms.MockAccounts, nil
}

func (ms *MockAccountStore) GetAccountByID(id int) (*models.AccountDTO, error) {
	for _, account := range MockAccounts {
		if account.ID == id {
			accountCopy := *account
			return &accountCopy, nil
		}
	}
	return nil, fmt.Errorf("account with ID %d not found", id)
}

func (ms *MockAccountStore) DeleteAccount(id int) error {
	for i, account := range ms.MockAccounts {
		if account.ID == id {
			ms.MockAccounts = append(ms.MockAccounts[:i], ms.MockAccounts[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("account not found")
}


func (ms *MockAccountStore) CreateAccount(acc *models.Account) error {
	return fmt.Errorf("create acc")
}

func (ms *MockAccountStore) UpdateAccount(acc *models.UpdateAccountReq, id int) error {
	return fmt.Errorf("update acc")
}

var MockAccounts = []*models.AccountDTO{
	{ID: 1, FirstName: "John", LastName: "Doe", Email: "john@example.com", Number: 12345, Balance: 100, CreatedAt: time.Now()},
	{ID: 2, FirstName: "James", LastName: "Doe", Email: "john@example.com", Number: 12345, Balance: 100, CreatedAt: time.Now()},
	{ID: 3, FirstName: "Mark", LastName: "Doe", Email: "john@example.com", Number: 12345, Balance: 100, CreatedAt: time.Now()},
}

func NewAccountMockStore() *MockAccountStore {
	return &MockAccountStore{MockAccounts: MockAccounts}
}
