## 使用方法

1. 直接运行
```sh
go run cli.go  --file model.go --dsn root:root@tcp\(localhost:3306\)/test?charset=utf8 --table user
```

2. 编译运行
```sh
go build -o t2t.bin cli.go
./t2t.bin -file model.go -dsn xxx -table user
```