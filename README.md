# 说明
本程序使用Go语言实现简单的电脑配置查询，用于记录电脑的CPU使用以及内存使用情况，同时支持汇总以及保存到CSV文件



## 代码依赖(gopm版本管理)

+ gopm get -g -v github.com/shirou/gopsutil/

> 如果需要打包Windows 版本需要另外安装如下依赖
+ gopm get -g -v github.com/StackExchange/wmi
+ gopm get -g -v github.com/go-ole/go-ole
+ gopm get -g -v github.com/go-ole/go-ole/oleutil

## 打包方式


+ MacOS 下打包

```shell script
$ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./src/Main.go
$ CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build ./src/Main.go
```


+ Linux 平台

```shell script
$ CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build ./src/Main.go
$ CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build ./src/Main.go
```

+ windows 平台

```shell script
$ SET CGO_ENABLED=0 SET GOOS=darwin3 SET GOARCH=amd64 go build ./src/Main.go
$ SET CGO_ENABLED=0 SET GOOS=linux SET GOARCH=amd64 go build ./src/Main.go
```