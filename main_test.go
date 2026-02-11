package main

import (
	"testing"

	"btpntest/domain"
	"btpntest/internal/cicilan/usecase"
)

func TestLoggingImports(t *testing.T) {
	mockRepo := &MockRepository{}
	useCase := usecase.NewCicilanUsecase(mockRepo)

	if useCase == nil {
		t.Fatal("Failed to create usecase")
	}
}

type MockRepository struct{}

func (m *MockRepository) GetAllTenors() ([]domain.Tenor, error) {
	return []domain.Tenor{
		{ID: 1, TenorValue: 6},
		{ID: 2, TenorValue: 12},
		{ID: 3, TenorValue: 24},
	}, nil
}

func TestValidationLogic(t *testing.T) {
	mockRepo := &MockRepository{}
	useCase := usecase.NewCicilanUsecase(mockRepo)

	req := &domain.CalculateInstallmentRequest{Amount: 10000000}
	resp, err := useCase.CalculateInstallments(req)

	if err != nil {
		t.Fatalf("Expected no error for valid amount, got %v", err)
	}

	if resp == nil {
		t.Fatal("Expected response, got nil")
	}

	if len(resp.Calculations) != 3 {
		t.Errorf("Expected 3 calculations, got %d", len(resp.Calculations))
	}
}

func TestRepositoryInitialization(t *testing.T) {
	mockRepo := &MockRepository{}

	tenors, err := mockRepo.GetAllTenors()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(tenors) != 3 {
		t.Errorf("Expected 3 tenors, got %d", len(tenors))
	}
}
