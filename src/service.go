package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Service struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func (s *Service) init() {
	s.initDB()
	s.initRouter()
}

func (s *Service) initDB() {
	db, err := gorm.Open("mysql", "root:990424@tcp(127.0.0.1:3306)/message4?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	fmt.Println("success connect to DB")
	s.DB = db

	db.SingularTable(true)

	if !s.DB.HasTable(&user{}) {
		s.DB.AutoMigrate(&user{})
		//s.DB.CreateTable(&user{})
		fmt.Println("create table user")
	}

	has := s.DB.HasTable(&Message{})
   	if !has {
     	s.DB.AutoMigrate(&Message{})
     	//db.CreateTable(&Message{})
     	fmt.Println("create table message")
   }

   s.DB.AutoMigrate(&user{}, &Message{})
   //添加列时自动修改表

}

func (s *Service) makeErrJSON(httpStatusCode int, errCode int, msg interface{}) (int, interface{}) {
	return httpStatusCode, map[string]interface{}{"error": errCode, "msg": fmt.Sprint(msg)}
}

func (s *Service) makeSuccessJSON(data interface{}) (int, interface{}) {
	return 200, map[string]interface{}{"error": 0, "msg": "success", "data": data}
}
