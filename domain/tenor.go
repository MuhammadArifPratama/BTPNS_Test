package domain

type Tenor struct {
	ID         int64 `gorm:"primaryKey"`
	TenorValue int   `gorm:"column:tenor_value;not null"`
	CreatedAt  int64 `gorm:"autoCreateTime:milli"`
	UpdatedAt  int64 `gorm:"autoUpdateTime:milli"`
}

func (Tenor) TableName() string {
	return "tenors"
}
