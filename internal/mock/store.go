package mock

import (
	"fmt"
	"time"

	"github.com/petar-dobre/gobank/internal/models"
)

type MockStore struct {
	MockAccounts []*models.AccountDTO
}

func (m *MockStore) GetAccounts() ([]*models.AccountDTO, error) {
	return m.MockAccounts, nil
}

func (m *MockStore) GetAccountByID(id int) (*models.AccountDTO, error) {
	for _, account := range MockAccounts {
		if account.ID == id {
			accountCopy := *account
			return &accountCopy, nil
		}
	}
	return nil, fmt.Errorf("account with ID %d not found", id)
}

func (m *MockStore) DeleteAccount(id int) error {
	for i, account := range m.MockAccounts {
		if account.ID == id {
			m.MockAccounts = append(m.MockAccounts[:i], m.MockAccounts[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("account not found")
}

func (s *MockStore) GetHashedPassword(email string) (string, error) {
	return "", nil
}

func (m *MockStore) CreateAccount(acc *models.Account) error {
	return fmt.Errorf("create acc")
}

func (m *MockStore) UpdateAccount(acc *models.Account) error {
	return fmt.Errorf("update acc")
}

var MockAccounts = []*models.AccountDTO{
	{ID: 1, FirstName: "John", LastName: "Doe", Email: "john@example.com", Number: 12345, Balance: 100, CreatedAt: time.Now()},
	{ID: 2, FirstName: "James", LastName: "Doe", Email: "john@example.com", Number: 12345, Balance: 100, CreatedAt: time.Now()},
	{ID: 3, FirstName: "Mark", LastName: "Doe", Email: "john@example.com", Number: 12345, Balance: 100, CreatedAt: time.Now()},
}

func NewMockStore() *MockStore {
	return &MockStore{MockAccounts: MockAccounts}
}
