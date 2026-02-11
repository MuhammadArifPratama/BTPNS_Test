package usecase

import (
	"testing"

	"btpntest/domain"
)

type MockCicilanRepository struct {
	tenors []domain.Tenor
	err    error
}

func (m *MockCicilanRepository) GetAllTenors() ([]domain.Tenor, error) {
	return m.tenors, m.err
}

func TestCalculateInstallments_Success(t *testing.T) {
	mockTenors := []domain.Tenor{
		{ID: 1, TenorValue: 6},
		{ID: 2, TenorValue: 12},
		{ID: 3, TenorValue: 24},
	}

	mockRepo := &MockCicilanRepository{tenors: mockTenors}
	usecase := NewCicilanUsecase(mockRepo)

	req := &domain.CalculateInstallmentRequest{Amount: 10000000}
	resp, err := usecase.CalculateInstallments(req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp == nil {
		t.Fatal("Expected response, got nil")
	}

	if len(resp.Calculations) != 3 {
		t.Errorf("Expected 3 calculations, got %d", len(resp.Calculations))
	}

	calc6 := resp.Calculations[0]
	if calc6.Tenor != 6 {
		t.Errorf("Expected tenor 6, got %d", calc6.Tenor)
	}
	if calc6.TotalMargin != 1000000 {
		t.Errorf("Expected total_margin 1000000, got %d", calc6.TotalMargin)
	}
	if calc6.TotalPayment != 11000000 {
		t.Errorf("Expected total_payment 11000000, got %d", calc6.TotalPayment)
	}
	if calc6.MonthlyInstallment != 1833333 {
		t.Errorf("Expected monthly_installment 1833333, got %d", calc6.MonthlyInstallment)
	}

	calc12 := resp.Calculations[1]
	if calc12.TotalMargin != 2000000 {
		t.Errorf("Expected total_margin 2000000, got %d", calc12.TotalMargin)
	}
	if calc12.TotalPayment != 12000000 {
		t.Errorf("Expected total_payment 12000000, got %d", calc12.TotalPayment)
	}

	calc24 := resp.Calculations[2]
	if calc24.TotalMargin != 4000000 {
		t.Errorf("Expected total_margin 4000000, got %d", calc24.TotalMargin)
	}
	if calc24.TotalPayment != 14000000 {
		t.Errorf("Expected total_payment 14000000, got %d", calc24.TotalPayment)
	}
}

func TestCalculateInstallments_InvalidAmount(t *testing.T) {
	mockRepo := &MockCicilanRepository{tenors: []domain.Tenor{}}
	usecase := NewCicilanUsecase(mockRepo)

	req := &domain.CalculateInstallmentRequest{Amount: 0}
	resp, err := usecase.CalculateInstallments(req)

	if err == nil {
		t.Fatal("Expected validation error for zero amount")
	}
	if resp != nil {
		t.Errorf("Expected nil response, got %v", resp)
	}

	req = &domain.CalculateInstallmentRequest{Amount: -1000}
	resp, err = usecase.CalculateInstallments(req)

	if err == nil {
		t.Fatal("Expected validation error for negative amount")
	}
}

func TestCalculateInstallments_NoTenors(t *testing.T) {
	mockRepo := &MockCicilanRepository{tenors: []domain.Tenor{}}
	usecase := NewCicilanUsecase(mockRepo)

	req := &domain.CalculateInstallmentRequest{Amount: 10000000}
	resp, err := usecase.CalculateInstallments(req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(resp.Calculations) != 0 {
		t.Errorf("Expected 0 calculations, got %d", len(resp.Calculations))
	}
}

func TestCalculateInstallments_DifferentAmount(t *testing.T) {
	mockTenors := []domain.Tenor{
		{ID: 1, TenorValue: 12},
	}

	mockRepo := &MockCicilanRepository{tenors: mockTenors}
	usecase := NewCicilanUsecase(mockRepo)

	req := &domain.CalculateInstallmentRequest{Amount: 5000000}
	resp, err := usecase.CalculateInstallments(req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	calc := resp.Calculations[0]

	if calc.TotalMargin != 1000000 {
		t.Errorf("Expected total_margin 1000000, got %d", calc.TotalMargin)
	}
	if calc.TotalPayment != 6000000 {
		t.Errorf("Expected total_payment 6000000, got %d", calc.TotalPayment)
	}
	if calc.MonthlyInstallment != 500000 {
		t.Errorf("Expected monthly_installment 500000, got %d", calc.MonthlyInstallment)
	}
}
