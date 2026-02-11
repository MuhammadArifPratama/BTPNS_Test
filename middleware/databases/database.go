package databases

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type DatabaseType string

const (
	MySQL      DatabaseType = "mysql"
	PostgreSQL DatabaseType = "postgresql"
	SQLServer  DatabaseType = "sqlserver"
)

type Config struct {
	Type     DatabaseType
	Host     string
	Port     int
	User     string
	Password string
	Database string
	SSLMode  string
}

func Connect(config Config) (*gorm.DB, error) {
	var dsn string

	switch config.Type {
	case MySQL:
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.User, config.Password, config.Host, config.Port, config.Database)
		return connectMySQL(dsn)

	case PostgreSQL:
		dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			config.Host, config.Port, config.User, config.Password, config.Database, config.SSLMode)
		return connectPostgreSQL(dsn)

	case SQLServer:
		dsn = fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
			config.User, config.Password, config.Host, config.Port, config.Database)
		return connectSQLServer(dsn)

	default:
		return nil, fmt.Errorf("unsupported database type: %s", config.Type)
	}
}

func connectMySQL(dsn string) (*gorm.DB, error) {
	log.Println("Connecting to MySQL...")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %w", err)
	}
	log.Println("✓ Connected to MySQL successfully")
	return db, nil
}

func connectPostgreSQL(dsn string) (*gorm.DB, error) {
	log.Println("Connecting to PostgreSQL...")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}
	log.Println("✓ Connected to PostgreSQL successfully")
	return db, nil
}

func connectSQLServer(dsn string) (*gorm.DB, error) {
	log.Println("Connecting to SQL Server...")
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to SQL Server: %w", err)
	}
	log.Println("✓ Connected to SQL Server successfully")
	return db, nil
}
