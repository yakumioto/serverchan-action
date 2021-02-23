# ServerChan Actions (Server酱)

ServerChan Actions(Server酱) 微信消息通知插件

## 简介

基于 [ServerChan(Server酱)](http://sc.ftqq.com/3.version) 封装的微信消息通知插件

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
        uses: yakumioto/serverchan-action@v1
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
        uses: yakumioto/serverchan-action@v1
        with:
          key: ${{ secrets.sckey }}
          title: {消息标题}
          desp: {消息内容 支持MarkDown}
```