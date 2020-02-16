package router

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
	"github.com/ChenYuxin0721/homework-during-winter-vacation/model"
	"github.com/ChenYuxin0721/homework-during-winter-vacation/db"
	"github.com/ChenYuxin0721/homework-during-winter-vacation/func"
)

func Router() {
	router := gin.Default()
	v1 := router.Group("/student")
	{
		v1.GET("/get/:id", GetStuinfo)
		v1.POST("/class",ListClass)
	}
	v2 := router.Group("/teacher") 
	{
		v2.POST("/list", ListStu) 
	}
	v3 := router.Group("/administrator")
	{
		v3.POST("/search",queryStu)
		v3.POST("/family/create", Family)
		v3.POST("/createStu", CreateStu)
		v3.DELETE("/:id", DeleteStu)
		v3.DELETE("/:id", DeleteTeacher)
	}

	router.GET("/stu", InitPage)

	router.Run(":8080")
}
