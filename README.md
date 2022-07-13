# Go语言学习笔记

### 做什么
- 熟悉 Go 基础语法
- 熟练使用 Go 技术栈相关组件
- 独立构建 Go web 项目
- 了解 Go 常见性能问题和优化方式

### todo
0. golang 的底层原理
1. 熟悉 GORM 框架，操作一遍数据库
2. 熟悉使用 Viper，处理好配置

### 如何做
- 首先找到高质量的信息渠道，阅读信息，训练。阅读什么信息
    - Go 语言学习，语言特性，生命周期
    - 有那些常用的库，语言适用场景是什么，有那些框架。
    - 语言架构是什么样的？代码执行流程是什么？
    - 时间和空间的管理，线程是如何设计的，内存是如何处理的？
- 会员3.0 项目代码阅读
- Go 语言项目搭建方式
- Go 并发性能，与Java 语言对比

### 目录结构
- based 基础语法实践
- web 构建 web 项目demo

### 语言特性
- Go 是静态类型的语言
- Go 是一门命令式语言，对比 Java 是一门声明式语言
- Go 语言是一门需要编译才能运行的编程语言，通过编译器生成二进制机器码，包含二进制机器码的文件才能在目标机器上运行。
- 鸭子类型（英语：duck typing）在程序设计中是动态类型的一种风格。在这种风格中，一个对象有效的语义，不是由继承自特定的类或实现特定的接口，而是由"当前方法和属性的集合"决定。

### 编程规范
- [官方编程规范中文版](https://www.gonglin91.com/2018/03/30/go-code-review-comments/)

### 常用类库
- fmt

### 常规语法
- 语言基础语法，变量，常量，数组的定义
- 基础流程，顺序执行，循环执行，逻辑判断
- 内存分配方式，new 和 mark
- 函数定义方式格式
    - 函数返回多个值
    - 函数不定参数
    - 闭包，函数中的函数。函数的定义与执行是分开的。用来读取局部变量变量，和保存局部变量的内存。
- 作用域，方法首字母大写，可被其他包访问
- 结构体的定义
- 接口的定义
- 错误处理，Error 接口；Defer； Panic/Recover 暴力停止进程

### 线程部分
- 创建线程 go func
- Channel
    ch <- 2 //发送数值2给这个通道
    x:=<-ch //从通道里读取值，并把读取的值赋值给x变量
    <-ch //从通道里读取值，然后忽略
- 定时器

### 网络编程
- socket编程

### 包管理
go module是Go1.11版本之后官方推出的版本管理工具，并且从Go1.13版本开始，go module将是Go语言默认的依赖管理工具。
```shell script
# 设置国内代理
go env -w GOPROXY=https://goproxy.cn,direct

# 常用命令
go mod download    # 下载依赖的module到本地cache（默认为$GOPATH/pkg/mod目录）
go mod edit        # 编辑go.mod文件
go mod graph       # 打印模块依赖图
go mod init        # 初始化当前文件夹, 创建go.mod文件
go mod tidy        # 增加缺少的module，删除无用的module
go mod vendor      # 将依赖复制到vendor下
go mod verify      # 校验依赖
go mod why         # 解释为什么需要依赖
```

### 通过问题理解
- [Golang 常见面试题](http://wearygods.online/articles/2021/04/19/1618823886966.html)


## web 相关知识学习记录

### 工具库记录
- github.com/rs/zerolog
The zerolog package provides a fast and simple logger dedicated to JSON output.

- github.com/spf13/cobra
Cobra is both a library for creating powerful modern CLI applications as well as a program to generate applications and command files.

- github.com/gin-gonic/gin
Gin is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to httprouter. If you need performance and good productivity, you will love Gin.

- github.com/spf13/viper
Viper is a complete configuration solution for Go applications including 12-Factor apps. It is designed to work within an application, and can handle all types of configuration needs and formats.

- gorm.io/gorm
The fantastic ORM library for Golang, aims to be developer friendly.

## 疑问余留
- mapstructure 是什么 「Env      string       `mapstructure:"env"`」的作用是什么
- 「data interface{}」是什么
- 过一遍 gin 的文档


# 信息链接
## 书籍
- [Go语言圣经](https://docs.hacknode.org/gopl-zh/)
- [Go语言高级编程](https://chai2010.cn/advanced-go-programming-book/)
- [GO专家编程](https://books.studygolang.com/GoExpertProgramming/)
- [Go 语言设计与实现](https://draveness.me/golang/)