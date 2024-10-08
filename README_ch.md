<!-- PROJECT SHIELDS -->

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

<!-- PROJECT LOGO -->
<br />

<p align="center">
  <a href="https://github.com/learnselfs/wlog/">
    <img src="logo.png" alt="Logo"  height="80">
  </a>

<h3 align="center"></h3>
  <p align="center">
wlog是Go（golang）的结构化记录器，与标准库记录器完全兼容API。
支持定时任务日志轮转。
    <br />
    <a href="https://github.com//learnselfs/wlog"><strong>探索本项目的文档 »</strong></a>
    <br />
    <br />
    <a href="https://github.com//learnselfs/wlog">查看Demo</a>
    ·
    <a href="https://github.com//learnselfs/wlog/issues">报告Bug</a>
    ·
    <a href="https://github.com//learnselfs/wlog/issues">提出新特性</a>
  </p>

</p>

[English](./README.md) | 中文 
## 目录

- [上手指南](#上手指南)
    - [开发前的配置要求](#开发前的配置要求)
    - [安装步骤](#安装步骤)
- [部署](#部署)
- [使用到的框架](#使用到的框架)
- [贡献者](#贡献者)
    - [如何参与开源项目](#如何参与开源项目)
- [版本控制](#版本控制)
- [鸣谢](#鸣谢)

### 上手指南
1. 直接使用 Debug、Info、Warn、Error、Fatal、Panic 
```go
wlog.Debug("test message")
wlog.Info("test message")
wlog.Warn("test message")
wlog.Error("test message")
wlog.Fatal("test message")
wlog.Panic("test message")
```
2. 标准库类型使用 
````go
wlog.Print("test message")
wlog.Println("test message")
wlog.Printf("%s","test message")

wlog.Painc("test message")
wlog.Paincln("test message")
wlog.Paincf("%s","test message")
````
3. 自定义配置其他项数据
```go
log := wlog.New()
log.Console()
log.WithKeys("key1", "key2", "key3")
for i := 1; i < 100; i++ {
    log.InfoValues(i+1, i+2, i+3)
}
```
4. 自定义输出格式(默认文本格式)
```go
l := wlog.New()
l.Json() // l.Text()
l.Info("test message")
```
5. 日志轮转
```go
log := wlog.NewLogConfig(InfoLevel, false, 0, false, wlog.NewFileCycle("info", wlog.DayCycle, "* * * * * *"), wlog.NewFormatJson())
for i := 0; i < 1000000; i++ {
log.WithField("item", strconv.Itoa(i))
    log.Println(i)
}

```
6. caller 调用信息
- 打印调用信息
```go
l := wlog.New()
l.CallFrame()
l.Info("test message")
```
- 打印错误日志
```go
// a.log
err := errors.New("error test")
Err(err)

//b.log
l := wlog.New()
l.CallFrameDepth() // call FrameDepth depth +1 
func Err(err error){
    ir err!=nil {
        l.error()
    }
}
```
7. 输出
```stdout
level="error"	 time="2024-**-** 13:24:03"	field="value"	file="D:/**/go/src/testing/testing.go"	func="testing.tRunner"	key="value"	line="1595"	
```
```json
{"error":"this log level is info, not call debug()","field":"value","file":"D:/**/go/src/testing/testing.go","func":"testing.tRunner","key":"value","level":"error","line":1595,"time":"2024-**-** 13:30:47"}
```
###### 开发前的配置要求

1. go version 1.21.1 

###### **安装步骤**

1. `go get github.com/learnselfs/wlog` 
[github.com/learnselfs/wlog](https://pkg.go.dev/github.com/learnselfs/wlog)

### 贡献者

请阅读**CONTRIBUTING.md** 查阅为该项目做出贡献的开发者。

#### 如何参与开源项目

贡献使开源社区成为一个学习、激励和创造的绝佳场所。你所作的任何贡献都是**非常感谢**的。


1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



### 版本控制

该项目使用Git进行版本管理。您可以在repository参看当前可用版本。

*您也可以在贡献者名单中参看所有参与该项目的开发者。*

### 版权说明

该项目签署了MIT 授权许可，详情请参阅 [LICENSE](https://github.com//learnselfs/wlog/blob/master/LICENSE)

### 鸣谢


- [logrus](https://github.com/sirupsen/logrus)
- [Best_README_template](https://github.com/shaojintian/Best_README_template)
- [Choose an Open Source License](https://choosealicense.com)
- [GitHub Pages](https://pages.github.com)

<!-- links -->
[your-project-path]:/learnselfs/wlog
[contributors-shield]: https://img.shields.io/github/contributors/learnselfs/wlog.svg?style=flat-square
[contributors-url]: https://github.com//learnselfs/wlog/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks//learnselfs/wlog.svg?style=flat-square
[forks-url]: https://github.com/learnselfs/wlog/network/members
[stars-shield]: https://img.shields.io/github/stars//learnselfs/wlog.svg?style=flat-square
[stars-url]: https://github.com//learnselfs/wlog/stargazers
[issues-shield]: https://img.shields.io/github/issues/learnselfs/wlog.svg?style=flat-square
[issues-url]: https://img.shields.io/github/issues//learnselfs/wlog.svg
[license-shield]: https://img.shields.io/github/license//learnselfs/wlog.svg?style=flat-square
[license-url]: https://github.com/learnselfs/wlog/blob/master/LICENSE
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=flat-square&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/shaojintian