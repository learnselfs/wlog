
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
wlog is a structured logger for Go (golang), completely API compatible with the standard library logger.
Support scheduled task log rotation.
    <br />
    <a href="https://github.com//learnselfs/wlog"><strong> markdown »</strong></a>
    <br />
    <br />
    <a href="https://github.com//learnselfs/wlog">Demo</a>
    ·
    <a href="https://github.com//learnselfs/wlog/issues">Bug</a>
    ·
    <a href="https://github.com//learnselfs/wlog/issues">Issues</a>
  </p>

</p>

English | [中文](./README_ch.md)
### Guide
1. Use Debug、Info、Warn、Error、Fatal、Panic
```go
wlog.Debug("test message")
wlog.Info("test message")
wlog.Warn("test message")
wlog.Error("test message")
wlog.Fatal("test message")
wlog.Panic("test message")
```
2. library log 
````go
wlog.Print("test message")
wlog.Println("test message")
wlog.Printf("%s","test message")

wlog.Painc("test message")
wlog.Paincln("test message")
wlog.Paincf("%s","test message")
````
3. Custom configuration of other item data
```go
log := wlog.New()
log.Console()
log.WithKeys("a", "b", "c")
for i := 1; i < 100; i++ {
    log.Values(i, i, i).Info("")
}
```
4. Custom output format (default text format)
```go
l := wlog.New()
l.Json() // l.Text()
l.Info("test message")
```
5. cron rotate log
```go
log := wlog.NewLogConfig(InfoLevel, false, 0, false, wlog.NewFileCycle("info", wlog.DayCycle, "* * * * * *"), wlog.NewFormatJson())
for i := 0; i < 1000000; i++ {
    log.WithField("item", strconv.Itoa(i))
    log.Println(i)
}
```
6. Caller Frame
- print caller Frame information
```go
l := wlog.New()
l.CallFrame()
l.Info("test message")
```
- print error information
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
7. Output 
```stdout
level="error"	 time="2024-**-** 13:24:03"	field="value"	file="D:/**/go/src/testing/testing.go"	func="testing.tRunner"	key="value"	line="1595"	
```
```json
{"error":"this log level is info, not call debug()","field":"value","file":"D:/**/go/src/testing/testing.go","func":"testing.tRunner","key":"value","level":"error","line":1595,"time":"2024-**-** 13:30:47"}
```
###### Pre development Configuration Requirements

1. go version 1.21.1

###### **Installation**

1. `go get github.com/learnselfs/wlog`
   [github.com/learnselfs/wlog](https://pkg.go.dev/github.com/learnselfs/wlog)

### Contributor 
Please read **CONTABUTING.md** to find out the developers who have contributed to this project.


#### Open Source Projects

Contributing makes the open source community an excellent place to learn, motivate, and create.
Any contribution you make is greatly appreciated.


1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

### License

This project has signed an Apache license, please refer to for details.
[LICENSE](https://github.com/learnselfs/wlog/blob/master/LICENSE)

### Thanks


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
