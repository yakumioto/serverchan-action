# ServerChan Actions (Server酱)

ServerChan Actions(Server酱) 微信消息通知插件

## 简介

基于 [ServerChan(Server酱)](https://sct.ftqq.com) 封装的微信消息通知插件

## 升级说明

> 因为微信发布公告将在4月底下线模板消息，Server酱开发了以企业微信为主的多通道新版（ Turbo版 sct.ftqq.com ）。旧版将在4月后下线，请尽快完成配置的更新。

`v1` 版本将在 4 月底不可用，请根据 [企业微信应用消息配置说明](https://sct.ftqq.com/forward) 进行升级。

`v2` 版本在不改变当前结构的情况下进行了升级，某些字段可能产生歧义。

- text 对应 turbo 版本中的 title 。

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
        uses: yakumioto/serverchan-action@v2
        with:
          key: {SCKEY}
          text: {消息标题}
          desp: {消息内容 支持 MarkDown}
          openid: {选填。抄送给其他关注你测试号的 openid ，多个 openid 用 , 隔开。}
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
        uses: yakumioto/serverchan-action@v2
        with:
          key: ${{ secrets.sckey }}
          text: {消息标题}
          desp: {消息内容 支持 MarkDown}
          openid: {选填。抄送给其他关注你测试号的 openid ，多个 openid 用 , 隔开。}
```