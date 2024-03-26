a lib for golang , generate mysql table schema to golang struct  
-----
mysql表结构自动生成golang struct  

## github地址
[https://github.com/gohouse/converter](https://github.com/gohouse/converter)

## 安装
1. 直接下载可执行文件: [下载地址](https://github.com/gohouse/converter/releases)  
2. golang源码包: `go get github.com/gohouse/converter`

## 示例表结构
```sql
CREATE TABLE `prefix_user` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Email` varchar(32) NOT NULL DEFAULT '' COMMENT '邮箱',
  `Password` varchar(32) NOT NULL DEFAULT '' COMMENT '密码',
  `CreatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表'
```

## 命令行用法
1. 下载对应平台的可执行文件, [下载地址](https://github.com/gohouse/converter/releases)

2. 命令行执行
    ```sh
    # 文件名: table2struct-[$platform].[$version].[$suffix]
    ./table2struct-linux.v0.0.3.bin -file model.go -dsn xxx -table user
    ```

3. 参数说明
```sh
-dsn            string 数据库dsn配置
-enableJsonTag  bool 是否添加json的tag
-file           string 保存路径
-packageName    string 包名
-prefix         string 表前缀
-realNameMethod string 结构体对应的表名
-table          string 要迁移的表
-tagKey         string tag的key
```

## golang代码简单用法
```go
package main
import (
	"fmt"
	"github.com/gohouse/converter"
)
func main() {
	err := converter.NewTable2Struct().
		SavePath("/home/go/project/model/model.go").
		Dsn("root:root@tcp(localhost:3306)/test?charset=utf8").
		Run()
	fmt.Println(err)
}
```

## golang代码详细用法示例
```go
package main

import (
	"fmt"
	"github.com/gohouse/converter"
)

func main() {
	// 初始化
	t2t := converter.NewTable2Struct()
	// 个性化配置
	t2t.Config(&converter.T2tConfig{
		// 如果字段首字母本来就是大写, 就不添加tag, 默认false添加, true不添加
		RmTagIfUcFirsted: false,
		// tag的字段名字是否转换为小写, 如果本身有大写字母的话, 默认false不转
		TagToLower: false,
		// 字段首字母大写的同时, 是否要把其他字母转换为小写,默认false不转换
		UcFirstOnly: false,
		// 每个struct放入单独的文件,默认false,放入同一个文件,true 单文件
		SeperatFile: true,
	})
	// 开始迁移转换
	err := t2t.
		// 指定某个表,如果不指定,则默认全部表都迁移
		// Table("user").
		// 表前缀
		// Prefix("prefix_").
		// 是否添加json tag
		EnableJsonTag(true).
		// 生成struct的包名(默认为空的话, 则取名为: package model)
		PackageName("model").
		// tag字段的key值,默认是orm
		TagKey("orm").
		// 是否添加结构体方法获取表名
		RealNameMethod("TableName").
		
		// 生成的结构体保存路径
		SavePath("/Users/fizz/go/src/github.com/gohouse/gupiao/model").
		// 数据库dsn,这里可以使用 t2t.DB() 代替,参数为 *sql.DB 对象
		Dsn("root:root@tcp(localhost:3306)/test?charset=utf8").
		// 执行
		Run()
	
	fmt.Println(err)
}
```

result 
```go
package model

import "time"

type User struct {
	Id         int     `json:"Id" orm:"Id"`
	Email      string  `json:"Email" orm:"Email"`           // 邮箱
	Password   string  `json:"Password" orm:"Password"`     // 密码
	CreatedAt  string  `json:"CreatedAt" orm:"CreatedAt"`
}

func (*User) TableName() string {
	return "user"
}
```
