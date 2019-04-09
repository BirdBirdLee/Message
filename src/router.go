package main

import "github.com/gin-gonic/gin"

func (s *Service) initRouter() {
	r := gin.Default()
	//users
	{
		//GET /users
		r.GET("/users", func(c *gin.Context) {
			c.JSON(s.getUsers(c))
		})
	}

	//user
	{
		//GET /user/{id}
		r.GET("/user/id/:id", func(c *gin.Context) {
			c.JSON(s.getUserById(c))
		})

		r.GET("/user/phone/:phone", func(c *gin.Context) {
			c.JSON(s.getUserByPhone(c))
		})

		//POST /user
		r.POST("/user", func(c *gin.Context) {
			c.JSON(s.postUser(c))
		})

		//PUT /user
		r.PUT("/user", func(c *gin.Context) {
			c.JSON(s.putUser(c))
		})

		//DELETE /user
		r.DELETE("/user/id/:id", func(c *gin.Context) {
			c.JSON(s.deleteUserById(c))
		})

		r.DELETE("/user/phone/:phone", func(c *gin.Context) {
			c.JSON(s.deleteUserByPhone(c))
		})

		
	}

	//Messages
	{
		//GET /Messages
		r.GET("/messages", func(c *gin.Context) {
			c.JSON(s.getMessages(c))
		})
	}

	//Message
	{
		r.GET("/message/id/:id", func(c *gin.Context) {
			c.JSON(s.getMessageById(c))
		})

		r.GET("/message/uid/:uid", func(c *gin.Context) {
			c.JSON(s.getMessageByUID(c))
		})

		r.GET("/message/content/:content", func(c *gin.Context) {
			c.JSON(s.getMessageByContent(c))
		})

		r.POST("/message", func(c *gin.Context) {
			c.JSON(s.postMessage(c))
		})

		r.PUT("/message", func(c *gin.Context) {
			c.JSON(s.putMessage(c))
		})

		r.DELETE("/message/id/:id", func(c *gin.Context) {
			c.JSON(s.deleteMessageById(c))
		})

		r.DELETE("/message/uid/:uid", func(c *gin.Context) {
			c.JSON(s.deleteMessageByUID(c))
		})
	}
	s.Router = r
	err := s.Router.Run(":8080")
	panic(err)
}
