package repository

import (
	"testing"

	"gorm.io/gorm"
)

func TestNewCicilanRepository(t *testing.T) {
	var db *gorm.DB
	repo := NewCicilanRepository(db)

	if repo == nil {
		t.Fatal("Failed to create repository")
	}

	if repo.db != db {
		t.Error("Repository db field not properly set")
	}
}

func TestRepositoryInterface(t *testing.T) {
	var db *gorm.DB
	repo := NewCicilanRepository(db)

	if repo != nil {
		t.Log("Repository created successfully for interface test")
	}
}
