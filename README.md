## Go语言学习笔记
- 首先找到高质量的信息渠道，阅读信息，训练。阅读什么信息
    - Go 语言学习，语言特性，生命周期
    - 有那些常用的库，语言适用场景是什么，有那些框架。
    - 语言架构是什么样的？代码执行流程是什么？
    - 时间和空间的管理，线程是如何设计的，内存是如何处理的？
- 会员3.0 项目代码阅读
- Go 语言项目搭建方式
- Go 并发性能，与Java 语言对比

### 语言特性
- go是静态类型的语言

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