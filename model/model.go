package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
    "log"
    "time"
    "github.com/gin-gonic/gin"
)

type Student struct {
	gorm.Familyinfo
	gorm.Classmate
    Name     string    `form:"name"`
	Address  string    `form:"address"`
	ID       int       `gorm:"primary_key"`
	Telenum  int   
	QQ       int       
    Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

type Familyinfo struct {
	Id   int 
	name string
	telenum int
}

type Classmate struct {
	Name string
	Telenum int16
}

type Teacher struct {
	Id   int 
	Name string
}