# dorylus
A golang job server

!["dorylus"](http://7xlu17.com1.z0.glb.clouddn.com/565ac46dd7a0b96e7c6f5336bcd98f72.jpg)
#这是什么？

这是一套分布式的cronjob 管理系统，支持manager机器基于负载和进程运行状态手动&自动调节分配运行的worker。

#快速开始

1. 先安装[gopm](https://github.com/gpmgo/gopm)
2. `go get -u github.com/gpmgo/gopm`
3. 确保你的GOPATH设置的环境变量可以调用gopm。
4. `gopm build`
5. `./dorylus`

#安装

确保依赖，首先下载无闻大大的[gopm](https://github.com/gpmgo/gopm)包，执行
`gopm install`
然后 `gopm run main.go`  就能运行了。
当然你可以build

#使用说明

##配置
修改conf/{环境}.conf,这里以开发环境dev为例

数据库的链接配置：

`mysql.conn = "{yourusername}:{youpassword}@/webcron?charset=utf8"`
`mysql.prefix = "t_"`

服务开启的端口：

`tcp.port = ":8989"`

最大运行进程数：

`WorkPollSize = 10`


这个只是一个worker，属于dorylus的执行者，管理者是另外的一个项目[dorylus-queen](https://github.com/bannerchi/dorylus-queen)。


#进度
1. 2016-03-15 Alpha设计完毕
2. 2016-03-18 worker框架代码提交
3. 2016-03-23 worker 完善
4. 2016-03-27 worker unit test
5. 2016-04-11 document


# 基于的项目
1. [beego的config](https://github.com/astaxie/beego)
2. [gotcp](https://github.com/gansidui/gotcp)
3. [webcron的jobs模块](https://github.com/lisijie/webcron)
4. [cron](https://github.com/robfig/cron)