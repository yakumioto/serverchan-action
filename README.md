# ServerChan Turbo Actions (Server酱)

ServerChan Actions(Server酱) 微信消息通知插件

## Fork 改动
> 因为微信发布公告将在4月底下线模板消息，Server酱开发了以企业微信为主的多通道新版（ Turbo版 sct.ftqq.com ）。旧版将在4月后下线，请尽快完成配置的更新。

[升级说明](http://sc.ftqq.com/9.version)


## 简介

基于 [ServerChan(Server酱)](https://sct.ftqq.com/) 封装的企业微信消息通知插件

## 栗子

明文 key 配置, 适用于私有仓库

```yaml
---
name: serverchan-action

on: [push]

jobs:
  serverchan:
    name: Server chan
    runs-on: ubuntu-latest
    steps:
      - name: Sending message
        uses: peixin/serverchan-action@v2.1
        with:
          key: {SCKEY}
          title: {消息标题}
          desp: {消息内容 支持MarkDown}
```

密文 key 配置, 适用于公共仓库.

```yaml
---
name: serverchan-action

on: [push]

jobs:
  serverchan:
    name: Server chan
    runs-on: ubuntu-latest
    steps:
      - name: Sending message
        uses: peixin/serverchan-action@v2.1
        with:
          key: ${{ secrets.sckey }}
          title: {消息标题}
          desp: {消息内容 支持MarkDown}
```
