package migration

import (
	"fmt"

	"gorm.io/gorm"
)

func RunMigrationForMySQL(db *gorm.DB) error {
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
	if err := db.Exec(sql).Error; err != nil {
		return err
	}

	seedSQL := `
	INSERT IGNORE INTO tenors (tenor_value, created_at, updated_at) VALUES
	(6, 0, 0),
	(12, 0, 0),
	(18, 0, 0),
	(24, 0, 0),
	(30, 0, 0),
	(36, 0, 0);
	`
	return db.Exec(seedSQL).Error
}

func RunMigrationForPostgreSQL(db *gorm.DB) error {
	if db.Migrator().HasTable(&Tenor{}) {
		return nil
	}

	sql := `
	CREATE TABLE IF NOT EXISTS tenors (
		id BIGSERIAL PRIMARY KEY,
		tenor_value INT NOT NULL UNIQUE,
		created_at BIGINT DEFAULT 0,
		updated_at BIGINT DEFAULT 0
	);
	`
	if err := db.Exec(sql).Error; err != nil {
		return err
	}

	seedSQL := `
	INSERT INTO tenors (tenor_value, created_at, updated_at) VALUES
	(6, 0, 0),
	(12, 0, 0),
	(18, 0, 0),
	(24, 0, 0),
	(30, 0, 0),
	(36, 0, 0)
	ON CONFLICT (tenor_value) DO NOTHING;
	`
	return db.Exec(seedSQL).Error
}

func RunMigrationForSQLServer(db *gorm.DB) error {
	if db.Migrator().HasTable(&Tenor{}) {
		return nil
	}

	sql := `
	IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='tenors' AND xtype='U')
	CREATE TABLE tenors (
		id BIGINT PRIMARY KEY IDENTITY(1,1),
		tenor_value INT NOT NULL UNIQUE,
		created_at BIGINT DEFAULT 0,
		updated_at BIGINT DEFAULT 0
	);
	`
	if err := db.Exec(sql).Error; err != nil {
		return err
	}

	seedSQL := `
	MERGE INTO tenors AS target
	USING (VALUES (6), (12), (18), (24), (30), (36)) AS source(tenor_value)
	ON target.tenor_value = source.tenor_value
	WHEN NOT MATCHED THEN
		INSERT (tenor_value, created_at, updated_at)
		VALUES (source.tenor_value, 0, 0);
	`
	return db.Exec(seedSQL).Error
}

func DetectDatabaseType(db *gorm.DB) string {
	if db.Dialector.Name() == "mysql" {
		return "mysql"
	} else if db.Dialector.Name() == "postgres" {
		return "postgres"
	} else if db.Dialector.Name() == "sqlserver" {
		return "sqlserver"
	}
	return "unknown"
}

func RunMigrationAuto(db *gorm.DB) error {
	dbType := DetectDatabaseType(db)

	switch dbType {
	case "mysql":
		return RunMigrationForMySQL(db)
	case "postgres":
		return RunMigrationForPostgreSQL(db)
	case "sqlserver":
		return RunMigrationForSQLServer(db)
	default:
		return fmt.Errorf("unsupported database type: %s", dbType)
	}
}
