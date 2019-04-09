package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
	"fmt"
)

type user struct {
	gorm.Model
	Username string
	Phone    string
	Password string
	Age      int
	Rank	 int8
}

type Message struct{
   gorm.Model
   //MID           int64    `gorm:"cloumn:MMID"`  
   Content       string   `gorm:"cloumn:MContent"`   
   //Valid         bool     `gorm:"cloumn:MValid"`           
   UID           int64   `gorm:"cloumn:UUID"`
}

func (s *Service) getUsers(c *gin.Context) (int, interface{}) {
	list := make([]*user, 0, 100)
	s.DB.Where(&user{}).Find(&list)
	return s.makeSuccessJSON(list)
}

func (s *Service) getUserById(c *gin.Context) (int, interface{}) {
	u := new(user)
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return s.makeErrJSON(400, 40000, "bad param")
	}
	s.DB.Where(&user{Model: gorm.Model{ID: uint(ID)}}).Find(u)
	if u.Phone == "" {
		return s.makeErrJSON(404, 40400, "not found")
	}
	return s.makeSuccessJSON(u)
}

func (s *Service) getUserByPhone(c *gin.Context) (int, interface{}) {
	u := new(user)
	phone := c.Param("phone")
	s.DB.Where("Phone = ?", phone).Find(u)
	if u.Phone == "" {
		return s.makeErrJSON(404, 40400, "not found")
	}
	return s.makeSuccessJSON(u)
}

func (s *Service) postUser(c *gin.Context) (int, interface{}) {
	u := new(user)
	err := c.BindJSON(u)
	if err != nil {
		return s.makeErrJSON(400, 40001, "bad payload")
	}
	u2 := new(user)
	s.DB.Where(&user{Phone: u.Phone}).Find(u2)
	if u2.Username != "" {
		return s.makeErrJSON(403, 40300, "phone already used")
	}
	tx := s.DB.Begin()
	if tx.Create(u).RowsAffected != 1 {
		tx.Rollback()
		return s.makeErrJSON(403, 50000, "create user error")
	}
	tx.Last(u2)
	tx.Commit()
	return s.makeSuccessJSON(map[string]interface{}{"uid": u2.ID})
}

func (s *Service) putUser(c *gin.Context) (int, interface{}) {
	u := new(user)
	err := c.BindJSON(u)
	if err != nil {
		return s.makeErrJSON(400, 40001, "bad payload")
	}
	u2 := new(user)
	s.DB.Where(&user{Model: gorm.Model{ID: u.ID}}).Find(u2)
	if u2.Username == "" {
		return s.makeErrJSON(404, 40400, "not found")
	}
	tx := s.DB.Begin()
	if tx.Model(&user{}).Where(&user{Model: gorm.Model{ID: u2.ID}}).Updates(u).RowsAffected != 1 {
		tx.Rollback()
		return s.makeErrJSON(403, 50000, "update user error")
	}
	tx.Commit()
	return s.makeSuccessJSON("success")
}

func (s *Service) deleteUserById(c *gin.Context) (int, interface{}) {
	u := new(user)
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return s.makeErrJSON(400, 40000, "bad param")
	}
	s.DB.Where(&user{Model: gorm.Model{ID: uint(ID)}}).Find(u)
	if u.Phone == "" {
		return s.makeErrJSON(404, 40400, "not found")
	}
	tx := s.DB.Begin()
	if tx.Model(&user{}).Where(&user{Model: gorm.Model{ID: uint(ID)}}).Delete(&user{Model: gorm.Model{ID: uint(ID)}}).RowsAffected != 1 {
		tx.Rollback()
		return s.makeErrJSON(403, 50000, "delete user error")
	}
	tx.Commit()
	return s.makeSuccessJSON("success")
}

func (s *Service) deleteUserByPhone(c *gin.Context) (int, interface{}) {
	u := new(user)
	phone := c.Param("phone")
	s.DB.Where("phone = ?", phone).Find(u)
	if u.Phone == "" {
		return s.makeErrJSON(404, 40400, "not found")
	}
	tx := s.DB.Begin()
	if tx.Where("phone = ?", phone).Delete(&user{}).RowsAffected != 1 {
		tx.Rollback()
		return s.makeErrJSON(403, 50000, "delete user error")
	}
	tx.Commit()
	return s.makeSuccessJSON("success")
}




func (s *Service) getMessages(c *gin.Context) (int, interface{}) {
	list := make([]*Message, 0, 100)
	s.DB.Where(&Message{}).Find(&list)
	return s.makeSuccessJSON(list)
}

func (s *Service) getMessageById(c *gin.Context) (int, interface{}) {
	m := new(Message)
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return s.makeErrJSON(400, 40000, "bad param")
	}
	s.DB.Where(&Message{Model: gorm.Model{ID: uint(ID)}}).Find(m)
	if m.Content == "" {
		return s.makeErrJSON(404, 40400, "not found")
	}
	return s.makeSuccessJSON(m)
}

func (s *Service) getMessageByUID(c *gin.Context) (int, interface{}) {
	list := make([]*Message, 0, 100)
	UID, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		return s.makeErrJSON(400, 40000, "bad param")
	}
	s.DB.Where("UID = ?", UID).Find(&list)
	return s.makeSuccessJSON(list)
}

func (s *Service) getMessageByContent(c *gin.Context) (int, interface{}) {
	list := make([]*Message, 0, 100)
	Content := c.Param("content")
	sql := fmt.Sprintf("select * from message where content like '%%%s%%'", Content)
	fmt.Println(sql)
	s.DB.Raw(sql).Scan(&list)
	//利用%内容%匹配
	return s.makeSuccessJSON(list)
}

func (s *Service) postMessage(c *gin.Context) (int, interface{}) {
	m := new(Message)
	err := c.BindJSON(m)
	if err != nil {
		return s.makeErrJSON(400, 40001, "bad payload")
	}
	/*
	m2 := new(Message)
	s.DB.Where(&user{Phone: u.Phone}).Find(u2)
	if u2.Username != "" {
		return s.makeErrJSON(403, 40300, "phone already used")
	}
	*/
	tx := s.DB.Begin()
	if tx.Create(m).RowsAffected != 1 {
		tx.Rollback()
		return s.makeErrJSON(403, 50000, "create message error")
	}
	tx.Create(m)
	tx.Commit()
	return s.makeSuccessJSON(map[string]interface{}{"mid": m.ID})
}

func (s *Service) putMessage(c *gin.Context) (int, interface{}) {
	//真实情况put的时候需要判断message的UID是否与当前用户的UID相同
	//管理员也不能修改
	m := new(Message)
	err := c.BindJSON(m)
	if err != nil {
		return s.makeErrJSON(400, 40001, "bad payload")
	}
	m2 := new(Message)
	s.DB.Where(&Message{Model: gorm.Model{ID: m.ID}}).Find(m2)
	if m2.Content == "" {
		return s.makeErrJSON(404, 40400, "not found")
	}
	tx := s.DB.Begin()
	if tx.Model(&Message{}).Where(&Message{Model: gorm.Model{ID: m2.ID}}).Updates(m).RowsAffected != 1 {
		tx.Rollback()
		return s.makeErrJSON(403, 50000, "update message error")
	}
	tx.Commit()
	return s.makeSuccessJSON("success")
}

func (s *Service) deleteMessageById(c *gin.Context) (int, interface{}) {
	//需要判断UID一致性
	//或Rank>1,即管理员
	m := new(Message)
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return s.makeErrJSON(400, 40000, "bad param")
	}
	s.DB.Where(&Message{Model: gorm.Model{ID: uint(ID)}}).Find(m)
	if m.Content == "" {
		return s.makeErrJSON(404, 40400, "not found")
	}
	tx := s.DB.Begin()
	if tx.Model(&Message{}).Where(&Message{Model: gorm.Model{ID: uint(ID)}}).Delete(&Message{Model: gorm.Model{ID: uint(ID)}}).RowsAffected != 1 {
		tx.Rollback()
		return s.makeErrJSON(403, 50000, "delete message error")
	}
	tx.Commit()
	return s.makeSuccessJSON("success")
}

func (s *Service) deleteMessageByUID(c *gin.Context) (int, interface{}) {
	list := make([]*Message, 0, 100)
	UID := c.Param("uid")
	s.DB.Where("uid = ?", UID).Find(&list)
	if len(list) == 0 {
		return s.makeErrJSON(404, 40400, "not found")
	}
	tx := s.DB.Begin()
	if tx.Where("uid = ?", UID).Delete(&Message{}).RowsAffected != int64(len(list)) {
		//删除的个数要和查询得到的个数一样
		tx.Rollback()
		return s.makeErrJSON(403, 50000, "delete message error")
	}
	tx.Commit()
	return s.makeSuccessJSON("success")
}
