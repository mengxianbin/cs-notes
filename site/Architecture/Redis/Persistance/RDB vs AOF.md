[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Redis](https://mengxianbin.github.io/cs-notes/site/Architecture/Redis) /
[Persistance](https://mengxianbin.github.io/cs-notes/site/Architecture/Redis/Persistance) /
[RDB vs AOF](https://mengxianbin.github.io/cs-notes/site/Architecture/Redis/Persistance/RDB%20vs%20AOF)

## RDB: Redis Database

> 全量模式

* 耗时
    * O(n)
    * fork 耗内存 (copy on write)
    * Disk IO 性能

* 不可控，容易丢失数据
    * 宕机

---

## AOF: Append Only File

> 增量模式，类似于 MySQL 的 binlog

---

| command    | RDB    | AOF    |
|------------|--------|--------|
| 启动优先级 | 低     | 高     |
| 体积       | 小     | 大     |
| 恢复速度   | 快     | 慢     |
| 数据安全性 | 丢数据 | 依策略 |
| 轻重       | 重     | 轻     |

---

## 最佳策略

### RDB

* 主关，从开

### AOF

* everysec

---
