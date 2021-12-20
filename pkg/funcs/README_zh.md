# 函数

## 全局默认函数

函数名称: `default`

特殊的，如果没有任何函数名能够匹配，则使用该函数。

### 参数

无需参数

### 表达式

`y = 0`

---

## 线性函数

Function name: `StandardLinearFunction`

### 参数

- slope
- offsetX
- offsetY

### 表达式

`y = slope * (x + offsetX) + offsetY`

### 内置预设函数

`DefaultLinearFunction`

- slope: 1
- offsetX: 0
- offsetY: 0

`ReverseLinearFunction`

- slope: -1
- offsetX: 0
- offsetY: 0

---

## 随机函数

函数名 `StandardRandomFunction`

### 参数

- range
- seed
- base-point

### 表达式

rand 由 seed 作为种子生成

y = base-point + rand()

base-point <= y <= base-point + range(ceil)

## 线性峰值函数

函数名 `StandardLinearPeak`

### 参数

- range
- offsetX
- offsetY
- ratio

### 表达式

y = ratio * ((x + offsetX) % range) + offsetY

### 内置预设函数

`SecondLinearPeak`

- range: 1
- ratio: 1
- offsetX: 0
- offsetY: 0

`MinuteLinearPeak`

- range: 60
- ratio: 1
- offsetX: 0
- offsetY: 0

`HourLinearPeak`

- range: 3600
- ratio: 1
- offsetX: 0
- offsetY: 0

`DayLinearPeak`

- range: 86400
- ratio: 1
- offsetX: 0
- offsetY: 0