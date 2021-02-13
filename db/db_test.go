package db

import (
	"testing"
	"fmt"
)

func TestGetDB(t *testing.T) {
	db, err := GetDB()

	fmt.Println(db)
	if err != nil {
		t.Fatal("failed test")
	}
}
