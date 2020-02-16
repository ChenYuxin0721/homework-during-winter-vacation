package main
import (
"database/sql"
"fmt"
"time"
_ "github.com/go-sql-driver/mysql"
)

type User struct {
ID int64 `db:"id"`

Name sql.NullString `db:"name"`
Age  int            `db:"age"`
}
// 定义mysql账号密码等信息
const (
USERNAME = "root"
PASSWORD = "12345678"
NETWORK  = "tcp"
SERVER  = "localhost"
PORT    = 3306
DATABASE = "tfbuilder"
)


func main() {
dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
DB, err := sql.Open("mysql", dsn)
if err != nil {
fmt.Printf("Open mysql failed,err:%v\n", err)
return
}
// 给db设置一个超时时间，时间小于数据库的超时时间即可
DB.SetConnMaxLifetime(100 * time.Second)
// 用于设置最大打开的连接数，默认值为0表示不限制。
DB.SetMaxOpenConns(100)
// 用于设置闲置的连接数
DB.SetMaxIdleConns(16)
queryOne(DB)
queryMulti(DB)
insertData(DB)
updateData(DB)
deleteData(DB)
}
//查询单行
func queryOne(DB *sql.DB) {
user := new(User)
// 查询单行数据
row := DB.QueryRow("select * from users where id=?", 1)
// 如果行不存在,则Scan()返回错误，需要处理异常，并绑定数据到结构体上
if err := row.Scan(&user.ID, &user.Name, &user.Age); err != nil {
fmt.Printf("scan failed, err:%v", err)
return
}
fmt.Println(*user)
}
//查询多行
func queryMulti(DB *sql.DB) {
user := new(User)
// 查询多行数据
rows, err := DB.Query("select * from users where id > ?", 1)
defer func() {
if rows != nil {
rows.Close()
}
}()
if err != nil {
fmt.Printf("Query failed,err:%v", err)
return
}
for rows.Next() {
err = rows.Scan(&user.ID, &user.Name, &user.Age)
if err != nil {
fmt.Printf("Scan failed,err:%v", err)
return
}
fmt.Print(*user)
}
}


//插入数据
func insertData(DB *sql.DB) {
result, err := DB.Exec("insert INTO users(name,age) values(?,?)", "YDZ", 23)
if err != nil {
fmt.Printf("Insert failed,err:%v", err)
return
}
// 最后插入的Id
lastInsertID, err := result.LastInsertId()
if err != nil {
fmt.Printf("Get lastInsertID failed,err:%v", err)
return
}
fmt.Println("LastInsertID:", lastInsertID)
// 本次插入数据影响的行数
rowsaffected, err := result.RowsAffected()
if err != nil {
fmt.Printf("Get RowsAffected failed,err:%v", err)
return
}
fmt.Println("RowsAffected:", rowsaffected)
}


func updateData(DB *sql.DB) {
result, err := DB.Exec("UPDATE users set age=? where id=?", "30", 3)
if err != nil {
fmt.Printf("Insert failed,err:%v", err)
return
}
rowsaffected, err := result.RowsAffected()
if err != nil {
fmt.Printf("Get RowsAffected failed,err:%v", err)
return
}
fmt.Println("RowsAffected:", rowsaffected)
}


func deleteData(DB *sql.DB) {
result, err := DB.Exec("delete from users where id=?", 1)
if err != nil {
fmt.Printf("Insert failed,err:%v", err)
return
}
rowsaffected, err := result.RowsAffected()
if err != nil {
fmt.Printf("Get RowsAffected failed,err:%v", err)
return
}
fmt.Println("RowsAffected:", rowsaffected)
}