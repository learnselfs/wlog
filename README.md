# wlog

#### Description
simple logger

#### Installation
`go get github.com/learnselfs/wlog`

#### Instructions

```go
l := wlog.New()
// l.SetFormatText() set text format output  
l.Info("data")
```
```json
{"level":1,"message":"msg","out":"out","time":"2024-02-26T10:52:43.6462286+08:00"}
```

```go
l := wlog.New()
f := make(Fields)
f["out"] = "out"
f["in"] = "in"
l.WithFields(f).Info("msg")
```
```json
{"in":"in","level":1,"message":"msg","out":"out","time":"2024-02-26T10:52:43.6462286+08:00"}
```

#### Contribution

1.  Fork the repository
2.  Create Feat_xxx branch
3.  Commit your code
4.  Create Pull Request


#### THANKS
- [Logrus](https://github.com/sirupsen/logrus)
