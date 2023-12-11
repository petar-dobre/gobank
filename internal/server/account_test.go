package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/petar-dobre/gobank/internal/mock"
	"github.com/petar-dobre/gobank/internal/models"
)

func compareAccountDTO(a, b *models.AccountDTO) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.ID == b.ID && a.FirstName == b.FirstName && a.LastName == b.LastName &&
		a.Email == b.Email && a.Number == b.Number && a.Balance == b.Balance &&
		a.CreatedAt.Equal(b.CreatedAt)
}

func TestHandleGetAccounts(t *testing.T) {
	mockStore := mock.NewMockStore()
	apiServer := &APIServer{store: mockStore}

	req, err := http.NewRequest("GET", "/accounts", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	err = apiServer.handleGetAccounts(rr, req)
	if err != nil {
		t.Fatalf("handleGetAccounts returned an unexpected error: %v", err)
	}

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var got []*models.AccountDTO
	if err := json.Unmarshal(rr.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}

	if len(got) != len(mock.MockAccounts) {
		t.Fatalf("expected %d accounts, got %d", len(mock.MockAccounts), len(got))
	}

	for i := range got {
		if !compareAccountDTO(got[i], mock.MockAccounts[i]) {
			t.Errorf("expected account %d to be %+v, got %+v", i, mock.MockAccounts[i], got[i])
		}
	}
}
