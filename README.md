# 小工具---自动生成虚拟币收益数据

执行exe文件在records目录下生成md文件，记录收益数据

## 配置说明
```json
{
  "records": [
    {
     "name": "bitcoin",
     "operate": "+",
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