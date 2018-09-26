# restday
中华人民共和国-节假日/休息日数据获取工具包

<br>

- 获取是否是工作日
```go
restday.Is_WorkDay("2018-06-18") // 端午节
restday.Is_WorkDay("2018-06-24") // 休息日
restday.Is_WorkDay("2018-06-25") // 工作日
```

- 获取月份天数
```go
restday.GetMonthDay("2018-06") // 30
```

- 获取月份休息日天数
```go
restday.GetMonthRestDay("2018-06") // 10
```

- 获取月份工作日天数
```go
restday.GetMonthWorkDay("2018-06") // 20
```

<br>

------