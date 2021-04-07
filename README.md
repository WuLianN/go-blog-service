### 开始使用

* 安装依赖，go install
* 执行程序，go run main.go

<br>

### 链路追踪

> 需要安装 Jaeger 
>
> https://www.jaegertracing.io/
>

<br>

windows 使用 Jaeger 

1. 解压下载的 Jaeger 压缩包
2. cd 到解压的目录，执行 jaeger-all-in-one.exe

<br>

### 将本地图片写入到数据库
```go
package main

import (
	"database/sql"
	"io/ioutil"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main(){
	var err error
	// db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/blog")
	db, err = sql.Open("mysql", "root:123456@/blog")

	err = db.Ping()
    if err != nil {
        panic(err.Error()) 
    }

	db.SetConnMaxLifetime(time.Minute * 3)
    db.SetMaxOpenConns(10)
    db.SetMaxIdleConns(10)

	defer db.Close()  

	const dirname = `C:\Users\smile 8\Desktop\beauty`
    
	files, _ := ioutil.ReadDir(dirname)

	sqlStr := "INSERT INTO blog_picture (name, file_name, state, is_del) VALUES "
	vals := []interface{}{}

	for _, file := range files {
		if !file.IsDir() {
			// 文件名 abc.jpg
		    fileName := file.Name()

		    // 去除后缀 abc 
			name := strings.Split(fileName, ".")[0]

			sqlStr += "(?, ?, ?, ?),"

			vals = append(vals, name, fileName, 1, 0)
		}
	}
	
	// trim the last ,
    sqlStr = sqlStr[0:len(sqlStr)-1]
 
    // prepare the statement
    stmt, _ := db.Prepare(sqlStr)

    // format all vals at once
    stmt.Exec(vals...)
}
``` 