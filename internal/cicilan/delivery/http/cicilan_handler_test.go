package http

import (
	"testing"

	"btpntest/domain"
)

// MockUsecase for testing the handler
type MockUsecase struct {
	response *domain.CalculateInstallmentResponse
	err      error
}

func (m *MockUsecase) CalculateInstallments(req *domain.CalculateInstallmentRequest) (*domain.CalculateInstallmentResponse, error) {
	return m.response, m.err
}

func TestNewCicilanHandler(t *testing.T) {
	mockUsecase := &MockUsecase{}
	handler := NewCicilanHandler(mockUsecase)

	if handler == nil {
		t.Fatal("Failed to create handler")
	}
}

func TestCicilanHandlerInit(t *testing.T) {
	mockResp := &domain.CalculateInstallmentResponse{
		Calculations: []domain.InstallmentCalculation{
			{
				Tenor:              6,
				MonthlyInstallment: 1833333,
				TotalMargin:        1000000,
				TotalPayment:       11000000,
			},
		},
	}

	mockUsecase := &MockUsecase{response: mockResp, err: nil}
	handler := NewCicilanHandler(mockUsecase)

	if handler.usecase == nil {
		t.Fatal("Handler usecase is nil")
	}
}
