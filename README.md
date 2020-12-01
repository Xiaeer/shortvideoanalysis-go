# shortvideoanalysis-go
短视频解析-go版

短视频解析demo（目前只有微视，后面有时间可能会分析其他）
基于Golang的第三方web框架 Gin 搭建的一个web服务

## Gin框架的详细教程看官方文档
```bash
https://gin-gonic.com/zh-cn/docs/
```

## 本地启动mongodb
首先要安装mongodb环境，因为短视频解析后的地址数据保存到mongodb，避免每次解析都要请求短视频分享链接导致效率问题

打开终端运行mongod
```bash
mongod
```

## 启动Gin服务
cd到此demo的根目录运行下面的启动命令
```bash
cd shortvideoanalysis

go run main.go
```

## 热重载启动Gin服务
使用bee工具（用于快速开发beego的工具，也可以用于Gin），需要自己go get安装
```bash
bee run
```

#### 短视频解析页面
```bash
http://127.0.0.1:8082/parseshortvideo
```
