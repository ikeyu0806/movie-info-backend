package db

import (
    "log"
    "io/ioutil"
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "gopkg.in/yaml.v2"
)

var (
  db  *gorm.DB
  err error
)

func GetDB() (*gorm.DB, error) {
  buf, err := ioutil.ReadFile("/workspace/golang_backend/db/db_connect.yml")
  if err != nil {
      log.Fatal(err)
  }

  var m map[string]interface{}

  err = yaml.Unmarshal([]byte(buf), &m)
  if err != nil {
      log.Fatal(err)
  }

  fmt.Println("!!!!!!!!!!!")
  fmt.Printf("%v\n", m)

  DBMS     := "mysql"
  USER     := "root"
  PASS     := "pass"
  PROTOCOL := "tcp(movie_mysql:3306)"
  DBNAME   := "movie_info"

  CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME+"?parseTime=true"
  db,err := gorm.Open(DBMS, CONNECT)

  if err != nil {
    panic(err.Error())
  }
  log.Println("connect mysql")
  return db, err
}
