# 小工具---自动生成虚拟币收益数据
## 下载
- [windows](https://github.com/K1ngram4/CoinRecord/releases/download/V1.0/CoinRecord_windows.zip)
- [osx](https://github.com/K1ngram4/CoinRecord/releases/download/V1.0/CoinRecord_mac.zip)

执行可执行文件即可在records目录下生成md文件，记录收益数据

- 效果

| 币种    | 持有数量 | 现价       | 总额         | 持仓成本     | 成本       | 利润      | 收益率 |
| ------- | -------- | ---------- | ------------ | ------------ | ---------- | --------- | ------ |
| bitcoin | 0.013294 | 45767.618385 | 608.445703 | 46171.875940 | 613.820000 | -5.374297 | -0.88% |
| dogecoin | 2927.600000 | 0.318420 | 932.206775 | 0.298500 | 694.863600 | 237.343175 | 34.16% |


## 配置说明
- 配置文件在holds目录下，json格式，多条记录分开
```json
{
  "records": [
    {
     "name": "bitcoin",
     "operate": "+",
     "amount": 0.01329424,
     "sum": 613.82
    },
    {
     "name": "bitcoin",
     "operate": "-",
     "amount": 0.01329424,
     "sum": 613.82
    }
  ]
}
```
- records 表示操作记录
- name 币种
- opetare 操作 “+”买入   “-”卖出
- amount 数量
- sum 总金额 单位usdt

**一个币种一个json文件**