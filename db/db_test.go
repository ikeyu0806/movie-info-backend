package db

import (
	"testing"
)

func TestGormConnect(t *testing.T) {
	db := gormConnect()
	if db == nil {
		t.Fatal("failed test")
	}
}

func TestGetDB(t *testing.T) {
	db := GetDB()

	if db == nil {
		t.Fatal("failed test")
	}
}
