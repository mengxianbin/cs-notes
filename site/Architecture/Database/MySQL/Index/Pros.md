[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Database](https://mengxianbin.github.io/cs-notes/site/Architecture/Database) /
[MySQL](https://mengxianbin.github.io/cs-notes/site/Architecture/Database/MySQL) /
[Index](https://mengxianbin.github.io/cs-notes/site/Architecture/Database/MySQL/Index) /
[Pros](https://mengxianbin.github.io/cs-notes/site/Architecture/Database/MySQL/Index/Pros)

# 索引的优点

* 减少需要扫描的数据行
* 避免排序、分组，避免创建临时表
    * B+ Tree 索引有序，用于 ORDER BY 和 GROUP BY 操作
    * 排序、分组过程会创建临时表

* 将随机 IO 变为顺序 IO
    * B+ Tree 索引有序，相邻数据会存储在一起

---

