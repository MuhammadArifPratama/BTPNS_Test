package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"btpntest/internal/cicilan/delivery/http"
	"btpntest/internal/cicilan/repository"
	"btpntest/internal/cicilan/usecase"
	"btpntest/internal/migration"
	"btpntest/middleware/databases"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("conf/conf.env"); err != nil {
		log.Printf("Warning: Could not load conf/conf.env: %v\n", err)
	}
}

func loadDatabaseConfig() databases.Config {
	dbType := os.Getenv("DB_TYPE")
	if dbType == "" {
		dbType = "mysql"
	}

	var dbTypeEnum databases.DatabaseType
	switch dbType {
	case "postgresql", "postgres", "pg":
		dbTypeEnum = databases.PostgreSQL
	case "sqlserver", "mssql", "sql-server":
		dbTypeEnum = databases.SQLServer
	default:
		dbTypeEnum = databases.MySQL
	}

	portStr := os.Getenv("DB_PORT")
	port := 3306
	if portStr != "" {
		if parsedPort, err := strconv.Atoi(portStr); err == nil {
			port = parsedPort
		}
	} else {
		switch dbTypeEnum {
		case databases.PostgreSQL:
			port = 5432
		case databases.SQLServer:
			port = 1433
		default:
			port = 3306
		}
	}

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		user = "root"
		switch dbTypeEnum {
		case databases.PostgreSQL:
			user = "postgres"
		case databases.SQLServer:
			user = "sa"
		}
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "password"
	}

	database := os.Getenv("DB_NAME")
	if database == "" {
		database = "btpntest"
	}

	sslMode := os.Getenv("DB_SSLMODE")
	if sslMode == "" {
		sslMode = "disable"
	}

	return databases.Config{
		Type:     dbTypeEnum,
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Database: database,
		SSLMode:  sslMode,
	}
}

func main() {
	dbConfig := loadDatabaseConfig()

	log.Printf("Connecting to database: %s at %s:%d\n", dbConfig.Type, dbConfig.Host, dbConfig.Port)

	db, err := databases.Connect(dbConfig)
	if err != nil {
		log.Printf("Error: Failed to connect to database: %v\n", err)
		log.Println("Application will continue, but database operations will fail.")
	} else {
		if err := migration.RunMigration(db); err != nil {
			log.Printf("Warning: Migration encountered an issue: %v\n", err)
			log.Println("Application will continue without migration.")
		}

		cicilanRepo := repository.NewCicilanRepository(db)
		cicilanUsecase := usecase.NewCicilanUsecase(cicilanRepo)
		cicilanHandler := http.NewCicilanHandler(cicilanUsecase)

		router := gin.Default()

		router.POST("/btpn/*path", func(c *gin.Context) {
			c.Request.URL.Path = c.Param("path")
			cicilanHandler.CalculateInstallments(c)
		})

		cicilanHandler.RegisterRoutes(router)

		server := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
		if server == ":" {
			server = ":8080"
		}

		log.Printf("Starting server on http://localhost%s\n", server)
		if err := router.Run(server); err != nil {
			log.Printf("Error: Failed to start server: %v\n", err)
		}
	}
}
