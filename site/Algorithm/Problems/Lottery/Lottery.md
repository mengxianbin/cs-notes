[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Algorithm](https://mengxianbin.github.io/cs-notes/site/Algorithm) /
[Problems](https://mengxianbin.github.io/cs-notes/site/Algorithm/Problems) /
[Lottery](https://mengxianbin.github.io/cs-notes/site/Algorithm/Problems/Lottery) /
[Lottery](https://mengxianbin.github.io/cs-notes/site/Algorithm/Problems/Lottery/Lottery)

# 概率抽奖

> https://www.cnblogs.com/zhoug2020/p/6396194.html

## 预生成奖品二维表

- N: 物品种类
- M: 概率最大值

| -      | -       |
| ------ | ------- |
| 预处理 | $O(MN)$ |
| 抽奖   | $O(1) $ |
| 空间   | $O(MN)$ |

## 预生成区间表

| -      | -                   |
| ------ | ------------------- |
| 预处理 | $O(N)             $ |
| 抽奖   | $O(N) \to O(log N)$ |
| 空间   | $O(N)             $ |


## Alias Method

> https://www.keithschwarz.com/darts-dice-coins/

| -      | -            |
| ------ | ------------ |
| 预处理 | $O(N log N)$ |
| 抽奖   | $O(1)$       |
| 空间   | $O(N)$       |

---
