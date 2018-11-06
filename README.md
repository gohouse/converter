a lib for golang , generate table schema to golang struct

## 用法示例
```go
package main

import (
	"fmt"
	"github.com/gohouse/converter"
)

func main() {
	t2t := converter.NewTable2Struct()
	err := t2t.
		Prefix("prefix_").
		RealNameMethod("TableName").
		SavePath("/path/to/model.go").
		Dsn("root:root@tcp(localhost:3306)/test?charset=utf8").
		Run()
	fmt.Println(err)
}
```

result 
```go
package model

import "time"

type User struct {
	Id         int       `json:"id"`
	Mobile     int       `json:"mobile"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	RealName   string    `json:"real_name"` // true name of user
	CreatedAt  time.Time `json:"created_at"`
}

func (t *User) TableName() string {
	return "prefix_user"
}
```
