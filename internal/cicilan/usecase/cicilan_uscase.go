package usecase

import (
	"btpntest/domain"
	cicilan "btpntest/internal/cicilan"
)

type cicilanUsecase struct {
	repo cicilan.CicilanRepository
}

func NewCicilanUsecase(repo cicilan.CicilanRepository) cicilan.CicilanUsecase {
	return &cicilanUsecase{repo: repo}
}

func (u *cicilanUsecase) CalculateInstallments(req *domain.CalculateInstallmentRequest) (*domain.CalculateInstallmentResponse, error) {

	if req.Amount <= 0 {
		return nil, &ValidationError{Message: "amount must be greater than 0"}
	}

	tenors, err := u.repo.GetAllTenors()
	if err != nil {
		return nil, err
	}

	if len(tenors) == 0 {
		return &domain.CalculateInstallmentResponse{
			Calculations: []domain.InstallmentCalculation{},
		}, nil
	}

	principal := req.Amount
	annualMarginRate := 0.2

	calculations := make([]domain.InstallmentCalculation, 0, len(tenors))

	for _, tenor := range tenors {
		totalMargin := int64(float64(principal) * annualMarginRate * (float64(tenor.TenorValue) / 12))

		totalPayment := principal + totalMargin

		monthlyInstallment := totalPayment / int64(tenor.TenorValue)

		calculations = append(calculations, domain.InstallmentCalculation{
			Tenor:              tenor.TenorValue,
			MonthlyInstallment: monthlyInstallment,
			TotalMargin:        totalMargin,
			TotalPayment:       totalPayment,
		})
	}

	return &domain.CalculateInstallmentResponse{
		Calculations: calculations,
	}, nil
}

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}
