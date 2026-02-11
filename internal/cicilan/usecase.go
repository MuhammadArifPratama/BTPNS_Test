package cicilan

import "btpntest/domain"

type CicilanUsecase interface {
	CalculateInstallments(req *domain.CalculateInstallmentRequest) (*domain.CalculateInstallmentResponse, error)
}
