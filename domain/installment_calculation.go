package domain

type CalculateInstallmentRequest struct {
	Amount int64 `json:"amount" binding:"required,gt=0"`
}

type InstallmentCalculation struct {
	Tenor              int   `json:"tenor"`
	MonthlyInstallment int64 `json:"monthly_installment"`
	TotalMargin        int64 `json:"total_margin"`
	TotalPayment       int64 `json:"total_payment"`
}

type CalculateInstallmentResponse struct {
	Calculations []InstallmentCalculation `json:"calculations"`
}
