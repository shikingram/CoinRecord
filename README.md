# 小工具---自动生成数字货币收益数据
## 下载
- [windows && osx](https://github.com/K1ngram4/CoinRecord/releases/download/v2.0/coinrecord_v1.2_template_modify.zip)

## 更新
*20210928更新：添加了设置代理访问
```shell
set https_proxy=http://127.0.0.1:1081
```

## 源码食用方式【windows环境】
- 效果==>[点击查看效果](http://kingram.top/posts/coin/20210928220055/)

- 1、下载golang环境[参考教程](http://kingram.top/posts/goland/goland_install/)
- 2、下载源码 `git clone https://github.com/K1ngram4/CoinRecord.git`
- 2、双击执行start.bat 即可在records目录下生成md文件，记录收益数据
- 

## 配置说明

- cfg下需要配置coinmarket.com的apikey，可以去申请[coinmarketcap.com/api](https://coinmarketcap.com/api/)

- 操作记录文件在holds目录下，json格式，多条记录分开（目前是手动添加，有时间做个网页服务）
```json
{
  "name": "btc",
  "records": [
    {
     "operate": "+",
     "amount": 0.01329424,
     "sum": 613.82
    },
    {
     "operate": "-",
     "amount": 0.01329424,
     "sum": 613.82
    }
  ]
}
```
- name 币种
- records 表示操作记录
  - opetare 操作 “+”买入   “-”卖出
  - amount 数量
  - sum 总金额 单位usdt

## 注意
- 1、一个币种一个json文件
- 2、秒退检查cmc的apikey是否填写以及是否科学上网
- 3、接口调不通请检查本地代理http端口是否正确