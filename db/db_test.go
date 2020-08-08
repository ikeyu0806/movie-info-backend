package db

import (
	"testing"
	"fmt"
)

func TestGormConnect(t *testing.T) {
	db, err := gormConnect()

	fmt.Println(db)
	if err != nil {
		t.Fatal("failed test")
	}
}

func TestGetDB(t *testing.T) {
	db, err := GetDB()

	fmt.Println(db)
	if err != nil {
		t.Fatal("failed test")
	}
}
