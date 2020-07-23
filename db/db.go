package db

import (
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

func gormConnect() *gorm.DB {
  DBMS     := "mysql"
  USER     := "root"
  PASS     := "pass"
  PROTOCOL := "tcp(movie_mysql:3306)"
  DBNAME   := "movie_info"

  CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
  db,err := gorm.Open(DBMS, CONNECT)

  if err != nil {
    panic(err.Error())
  }
  log.Println("connect mysql")
  return db
}

func Init() {
  db := gormConnect()
  defer db.Close()
} 

func GetDB() *gorm.DB {
  DBMS     := "mysql"
  USER     := "root"
  PASS     := "pass"
  PROTOCOL := "tcp(movie_mysql:3306)"
  DBNAME   := "movie_info"

  CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
  db,err := gorm.Open(DBMS, CONNECT)

  if err != nil {
    panic(err.Error())
  }
  log.Println("connect mysql")
  return db
}
