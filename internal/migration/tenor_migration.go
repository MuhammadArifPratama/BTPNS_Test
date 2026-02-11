package migration

import "gorm.io/gorm"

func RunMigration(db *gorm.DB) error {
	if db.Migrator().HasTable(&Tenor{}) {
		return nil
	}

	return RunMigrationAuto(db)
}

func CreateTenorTable(db *gorm.DB) error {
	if db.Migrator().HasTable(&Tenor{}) {
		return nil
	}

	sql := `
	CREATE TABLE IF NOT EXISTS tenors (
		id BIGINT AUTO_INCREMENT PRIMARY KEY,
		tenor_value INT NOT NULL,
		created_at BIGINT DEFAULT 0,
		updated_at BIGINT DEFAULT 0,
		UNIQUE KEY unique_tenor_value (tenor_value)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
	`
	return db.Exec(sql).Error
}

func SeedTenorData(db *gorm.DB) error {
	var count int64
	if err := db.Model(&Tenor{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	sql := `
	INSERT IGNORE INTO tenors (tenor_value, created_at, updated_at) VALUES
	(6, 0, 0),
	(12, 0, 0),
	(18, 0, 0),
	(24, 0, 0),
	(30, 0, 0),
	(36, 0, 0);
	`
	return db.Exec(sql).Error
}

type Tenor struct {
	ID         int64 `gorm:"primaryKey"`
	TenorValue int   `gorm:"column:tenor_value;not null"`
	CreatedAt  int64 `gorm:"autoCreateTime:milli"`
	UpdatedAt  int64 `gorm:"autoUpdateTime:milli"`
}

func (Tenor) TableName() string {
	return "tenors"
}
