# dorylus
A golang job server
#这是什么？

这是一套分布式的cronjob 管理系统，支持manager机器基于负载和进程运行状态手动&自动调节分配运行的woker


#进度
1. 2016-03-15 Alpha设计完毕
2. 2016-03-18 worker框架代码提交
3. 2016-03-23 woker 完善
4. 2016-03-27 worker unit test


# 基于的项目
1. [beego的config](https://github.com/astaxie/beego)
2. [gotcp](https://github.com/gansidui/gotcp)
3. [webcron的jobs模块](https://github.com/lisijie/webcron)
4. [cron](https://github.com/robfig/cron)