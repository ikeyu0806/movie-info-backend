package models

import (
	"fmt"
	"testing"
  "github.com/gin-gonic/gin"
	"net/http"
	"bytes"
)

func TestCreateReview(t *testing.T) {
	jsonStr := `{"public_id":123456789, "comment":"good", "score":12345}`
	req, _ := http.NewRequest("POST", "/review/create", bytes.NewBuffer([]byte(jsonStr)))

	var context *gin.Context
	context = &gin.Context{Request: req}
	
	var r ReviewModel
	result, err := r.CreateReview(context)

	fmt.Println(result)
	if err != nil {
		fmt.Println(err)
		t.Fatal(err)
	}
}

func TestFindByMovieId(t *testing.T) {
	req, _ := http.NewRequest("GET", "/reviews/1", nil)

	var context *gin.Context
	context = &gin.Context{Request: req}
	
	var r ReviewModel
	result, err := r.FindByMovieId(context)

	fmt.Println(result)
	if err != nil {
		fmt.Println(err)
		t.Fatal(err)
	}
}
