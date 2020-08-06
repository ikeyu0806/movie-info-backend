package models

import (
	"fmt"
	"testing"
  "github.com/gin-gonic/gin"
  "net/http"
)

func TestFindByMovieID(t *testing.T) {
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
