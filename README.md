# dorylus
[![Build Status](https://circleci.com/gh/bannerchi/dorylus.png?circle-token=d5c4d4e8578ad9a4f56038e807b192fc7091adef)](https://circleci.com/gh/bannerchi/dorylus)

A golang job server

!["dorylus"](http://7xlu17.com1.z0.glb.clouddn.com/565ac46dd7a0b96e7c6f5336bcd98f72.jpg)

# 这是什么？

这是一套分布式的cronjob 管理系统，支持manager机器基于负载和进程运行状态手动&自动调节分配运行的worker。

# 快速开始

1. 先安装[gopm](https://github.com/gpmgo/gopm)
2. `go get -u github.com/gpmgo/gopm`
3. 确保你的GOPATH设置的环境变量可以调用gopm。
4. `gopm build`
5. `./dorylus`

# 安装

确保依赖，首先下载无闻大大的[gopm](https://github.com/gpmgo/gopm)包，执行 <br>
`gopm install` <br>
然后 `gopm run main.go`  就能运行了。<br>
当然你也可以go build 生成执行文件 dorylus, 执行 ./dorylus ,<br>请确保你的配置文件已经配好。

# 使用说明

## 配置
修改conf/{环境}.conf,这里以开发环境dev为例

数据库的链接配置：

`mysql.conn = "{yourusername}:{youpassword}@/webcron?charset=utf8"`
`mysql.prefix = "t_"`

服务开启的端口：

`tcp.port = ":8989"`

最大运行进程数：

`WorkPollSize = 10`


这个只是一个worker，属于dorylus的执行者，管理者是另外的一个项目[dorylus-queen](https://github.com/bannerchi/dorylus-queen)。


# 进度
- [x] 2016-03-15 Alpha设计完毕
- [x] 2016-03-18 worker框架代码提交
- [x] 2016-03-23 worker 完善
- [x] 2016-03-27 worker unit test
- [x] 2016-04-07 document
- [x] 2016-11-07  [dorylus-cli](https://github.com/bannerchi/dorylus-cli)
- [ ] worker 单例和并用模式切换（new）
- [ ] 断线还原job进度（new）
- [ ] job探针


# 基于的项目
1. [beego的config](https://github.com/astaxie/beego)
2. [gotcp](https://github.com/gansidui/gotcp)
3. [webcron的jobs模块](https://github.com/lisijie/webcron)
4. [cron](https://github.com/robfig/cron)
