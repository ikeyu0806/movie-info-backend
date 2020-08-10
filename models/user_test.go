package models

import (
	"fmt"
	"testing"
  "github.com/gin-gonic/gin"
	"net/http"
	"bytes"
)

func TestCreateUser(t *testing.T) {
	jsonStr := `{"name":"testname", "email":"example.com", "password":"pass"}`
	req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer([]byte(jsonStr)))

	var context *gin.Context
	context = &gin.Context{Request: req}
	
	var r UserModel
	result, err := r.CreateUser(context)

	fmt.Println(result)
	if err != nil {
		fmt.Println(err)
		t.Fatal(err)
	}
}

func TestGetByName(t *testing.T) {
	name := "test"
	
	var r UserModel
	result, err := r.FindByName(name)

	fmt.Println(result)
	if err != nil {
		fmt.Println(err)
		t.Fatal(err)
	}
}
