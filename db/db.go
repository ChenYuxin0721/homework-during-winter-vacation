package main

import (
"database/sql"
"fmt"
"time"
_ "github.com/go-sql-driver/mysql"
"github.com/ChenYuxin0721/homework-during-winter-vacation/router"
"github.com/ChenYuxin0721/homework-during-winter-vacation/model"
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
}