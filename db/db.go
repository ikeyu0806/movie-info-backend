package db

import (
    "log"
    "os"
    "io/ioutil"
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

  env := os.Getenv("SERVICE_ENV")

  switch env {
  case "production":
    DBMS     := "mysql"
    USER     := m["prooduction"].(map[interface {}]interface {})["user"].(string)
    PASS     := m["prooduction"].(map[interface {}]interface {})["password"].(string)
    PROTOCOL := m["prooduction"].(map[interface {}]interface {})["protocol"].(string)
    DBNAME   := m["prooduction"].(map[interface {}]interface {})["db_name"].(string)
  
    CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME+"?parseTime=true"
    db,err := gorm.Open(DBMS, CONNECT)
  
    if err != nil {
      panic(err.Error())
    }
    log.Println("connect mysql")
    return db, err
  case "test":
    DBMS     := "mysql"
    USER     := m["test"].(map[interface {}]interface {})["user"].(string)
    PASS     := m["test"].(map[interface {}]interface {})["password"].(string)
    PROTOCOL := m["test"].(map[interface {}]interface {})["protocol"].(string)
    DBNAME   := m["test"].(map[interface {}]interface {})["db_name"].(string)
  
    CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME+"?parseTime=true"
    db,err := gorm.Open(DBMS, CONNECT)
  
    if err != nil {
      panic(err.Error())
    }
    log.Println("connect mysql")
    return db, err
  default:
    DBMS     := "mysql"
    USER     := m["development"].(map[interface {}]interface {})["user"].(string)
    PASS     := m["development"].(map[interface {}]interface {})["password"].(string)
    PROTOCOL := m["development"].(map[interface {}]interface {})["protocol"].(string)
    DBNAME   := m["development"].(map[interface {}]interface {})["db_name"].(string)
  
    CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME+"?parseTime=true"
    db,err := gorm.Open(DBMS, CONNECT)
  
    if err != nil {
      panic(err.Error())
    }
    log.Println("connect mysql")
    return db, err
  }
}
