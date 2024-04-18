## jiacrontab

[![Build Status](https://travis-ci.org/iwannay/jiacrontab.svg?branch=dev)](https://travis-ci.org/iwannay/jiacrontab) 

简单可信赖的任务管理工具

## 原仓库地址 https://github.com/iwannay/jiacrontab

### 说明
使用 go1.16+ embed 打包静态文件
兼容 mysql8.0+ 默认配置,使用Utf8mb4编码

```
// 跨平台编译设置
set GOOS=linux  
set GOARCH=amd64
set CGO_ENABLED=1

// 跨平台编译设置 windows powerShell
$env:GOOS="linux"
$env:GOARCH="amd64"
$env:CGO_ENABLED="1"

 go build  -o .\bin\jiacrontab_admin -a .\app\jiacrontab_admin\main.go
 go build  -o .\bin\jiacrontabd .\app\jiacrontabd\main.go
```
```

### v2.0.0版发布


### [❤jiacrontab 最新版下载点这里❤ ](https://download.iwannay.cn/jiacrontab/)

    1.自定义job执行  
    2.允许设置job的最大并发数  
    3.每个脚本都可在web界面下灵活配置，如测试脚本运行，查看日志，强杀进程，停止定时...  
    4.允许添加脚本依赖（支持跨服务器），依赖脚本提供同步和异步的执行模式  
    5.支持异常通知  
    6.支持守护脚本进程  
    7.支持节点分组


### 架构

<img src="https://raw.githubusercontent.com/iwannay/static_dir/master/jiacrontab_arch.png" width="50%"/>

### 说明

jiacrontab 由 jiacrontab_admin，jiacrontabd 两部分构成，两者完全独立通过 rpc 通信  
jiacrontab_admin：管理后台向用户提供web操作界面  
jiacrontabd：负责job数据存储，任务调度  


### 安装

#### 二进制安装

1.[下载](https://download.iwannay.cn/jiacrontab/) 二进制文件。

2.解压缩进入目录(jiarontab_admin,jiacrontabd)。

3.运行

```sh
$ nohup ./jiacrontab_admin &> jiacrontab_admin.log &
$ nohup ./jiacrontabd &> jiacrontabd.log &

## 建议使用systemd守护
```
