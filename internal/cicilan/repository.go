package cicilan

import "btpntest/domain"

type CicilanRepository interface {
	GetAllTenors() ([]domain.Tenor, error)
}
