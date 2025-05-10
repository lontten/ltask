# ltask - 定时任务管理系统

# 项目简介

使用Go语言开发的轻量级定时任务集中调度和管理系统

## 功能特性

* 两种执行逻辑，1直接任务执行，在web界面添加任务；2基于元数据的自定义任务执行器。
* Web界面管理定时任务
* crontab时间表达式, 精确到秒
* 任务执行失败可重试
* 任务执行超时, 强制结束
* 任务依赖配置, A任务完成后再执行B任务
* 账户权限控制
* 任务类型
    * shell任务
  > 在任务节点上执行shell命令, 支持任务同时在多个节点上运行
    * HTTP任务
  > 访问指定的URL地址, 由调度器直接执行, 不依赖任务节点
* 查看任务执行结果日志
* 任务执行结果通知, 支持邮件、Webhook

### 支持平台

> Windows、Linux、Mac OS

### 环境要求

> MySQL
> Postgresql

## 安装

### 二进制安装

1. 解压压缩包
2. `cd 解压目录`
3. 启动

* 调度器启动
    * Windows: `ltask.exe web`
    * Linux、Mac OS:  `./ltask web`
* 任务节点启动, 默认监听0.0.0.0:11171
    * Windows:  `ltask-node.exe`
    * Linux、Mac OS:  `./ltask-node`

4. 浏览器访问 http://localhost:11170

### 源码安装

- 安装Go 1.23+

### 开发

1. 安装Go1.23+, Node.js 20+, pnpm

### 命令

* ltask
    * -v 查看版本

## To Do List

- [] 添加任务
- [] web页面
- [] 版本升级
- [] 批量开启、关闭、删除任务
- [] 调度器与任务节点通信支持https
- [] 任务分组
- [] 多用户
- [] 权限控制

## 程序使用的组件

* 定时任务调度 [Cron](https://github.com/robfig/cron)
* UI框架 [element-plus](https://github.com/element-plus/element-plus)

## ChangeLog
