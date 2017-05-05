weixin
======

Golang 实现的微信公众号接口

### Install

```bash
go get -u -v github.com/arstd/weixin
```

### Test

```bash
cd $GOPATH/src/github.com/arstd/weixin
# go get ./... # 依赖于 arstd/log
go test ./... 执行所有单元测试
go run examples/*.go 启动一个 Server
```

### Usage

见单元测试和 examples

### 如何本地调试微信公众号？

[参考 Debug-weixin-locally.md](Debug-weixin-locally.md)


### 测试公众号

扫码关注如下公众号

![测试号二维码](http://mmbiz.qpic.cn/mmbiz/Ls7EibW7x9GmxYSNSibDAqeqCPJ7Axo2BmLyTrRPbZMhiaS7IfHBlmz0xiaNcAX9LdcsQBub8V6aibY2bEsw3iapAmlQ/0)


### TODO

* Token 单元测试没过
* iBeacon 等更多接口实现
* 多个公众号
