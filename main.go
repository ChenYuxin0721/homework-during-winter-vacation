package main

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


func main() {
    route := gin.Default()
    route.GET("/testing", startPage)
    route.Run(":8085")
}


var (
	db  *gorm.DB
	err error
)


func main() {
	db, err := sql.Open("mysql", "root:123456/nulige?charset=utf8")
	if err != nil {
		panic(err)
	}

	//fmt.Println(db.Ping())  检查是否连接成功数据库
	stmt, err := db.Prepare("INSERT INTO user_info SET username=?,departname=?,create_time=?")
	if err != nil {
		fmt.Println(err)
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	Router()
}

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

func GetStuinfo(c *gin.Context) {
	var stu Stu
	db.Create(&stu)
	id := c.Params.ByName("id")
	err := db.First(&stu, id).Error
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err.Error())
	} else {
		c.JSON(200, &stu)
	}
}

func ListStu(c *gin.Context) {
	var stu []Stu
	db.Find(&stu)
	c.JSON(200, &stu)
}


func ListClass(c *gin.Context) {
	var class []Class
	db.Find(&class)
	c.JSON(200, &class)
}


func DeleteStu(c *gin.Context) {
	var stu Stu
	id := c.Params.ByName("id")
	db.First(&stu, id)
	if user.Id != 0 {
		db.Delete(&stu)
		c.JSON(200, gin.H{
			"message": "successfully deleted",
		})
	} else {
		c.JSON(404, gin.H{
			"error": "Not found",
		})
	}
}


func DeleteTeacher(c *gin.Context) {
	var teacher Teacher
	id := c.Params.ByName("id")
	db.First(&teach, id)
	if teacher.Id != 0 {
		c.JSON(200, gin.H{
			"success": "successfully deleted",
		})
	} else {
		c.JSON(404, gin.H{
			"error": "Not found",
		})
	}
}


func Family(c *gin.Context) {
	var family Familyinfo
	c.BindJSON(&family)
	if family.ID != "" {
		db.Create(&family)
		c.JSON(200, gin.H{"success": &family})
	} else {
		c.JSON(400, gin.H{
			"message": "error",
		})
	}
}

func CreateStu(c *gin.Context) {
	var stu Stu
	c.BindJSON(&stu)
	if stu.ID != "" {
		db.Create(&stu)
		c.JSON(200, gin.H{"success": &stu})
	} else {
		c.JSON(400, gin.H{
			"message": "error",
		})
	}
}

func updateData(c *gin.Context) {
	result, err := DB.Exec("UPDATE users set age=? where id=?", "30", 3)
	if err != nil {
	fmt.Printf("Insert failed,err:%v", err)
	return
	}
	rowsaffected, err := result.RowsAffected()
	if err != nil {
	c.JSON(400,gin.H{
		"error":"NOT FIND"
	})
	return
	}
	c.JSON(200,rowsaffected)
	}

func queryStu(c *gin.Context) {
	user := new(Stu)
	row := DB.QueryRow("select * from users where id=?", 1)
	// 如果行不存在,则Scan()返回错误，需要处理异常，并绑定数据到结构体上
	if err := row.Scan(&stu.ID, &stu.Name, &stu.telenum); err != nil {
	c.JSON(400,gin.H{
		"error":"NOT FIND"
	})
	return
	}
	c.JSON(200,&stu)
	}