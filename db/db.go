package db

import (
    "os"
    "log"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
  db  *gorm.DB
  err error
)

type User struct {
	ID        uint
	Name      string
	Password  string
}

func GetDB() (*gorm.DB, error) {
  DBMS     := "mysql"
  USER     := os.Getenv("DB_USER")
  PASS     := os.Getenv("DB_PASSWORD")
  PROTOCOL := "tcp(" + os.Getenv("DB_HOST") + ")"
  DBNAME   := os.Getenv("DB_NAME")

  CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME+"?parseTime=true"
  db,err := gorm.Open(DBMS, CONNECT)

  if err != nil {
    panic(err.Error())
  }
  log.Println("connect mysql")
  return db, err
}
