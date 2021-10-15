# 小工具---自动生成虚拟币收益数据
## 下载
- [windows && osx](https://github.com/K1ngram4/CoinRecord/releases/download/v2.0/coinrecord_v1.2_template_modify.zip)

- 20210928更新：添加了设置代理访问

```shell
set https_proxy=http://127.0.0.1:1081
```

执行可执行文件即可在records目录下生成md文件，记录收益数据

- 效果

[点击查看效果](http://kingram.top/posts/coin/20210928220055/)

## 配置说明

- cfg下需要配置coinmarket.com的apikey，可以去申请[coinmarketcap.com/api](https://coinmarketcap.com/api/)

- 操作配置文件在holds目录下，json格式，多条记录分开
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
  
**一个币种一个json文件**