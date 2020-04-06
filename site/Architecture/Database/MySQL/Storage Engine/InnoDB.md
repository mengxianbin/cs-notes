[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Database](https://mengxianbin.github.io/cs-notes/site/Architecture/Database) /
[MySQL](https://mengxianbin.github.io/cs-notes/site/Architecture/Database/MySQL) /
[Storage Engine](https://mengxianbin.github.io/cs-notes/site/Architecture/Database/MySQL/Storage%20Engine) /
[InnoDB](https://mengxianbin.github.io/cs-notes/site/Architecture/Database/MySQL/Storage%20Engine/InnoDB)

# InnoDB

* 支持事务
* 默认隔离级别：可重复读
* 主索引使聚簇索引
    * 索引中保持数据，避免直接读取磁盘，对性能有很大提升

* 支持在线热备份
    * 获取一致性视图不需要禁止读写

* 实现读诸多优化
    * 可预测性读
    * 自适应哈希索引
        * 自动创建
        * 加快读取
    * 插入缓冲区
        * 快速插入

* 支持外键

## 隔离级别

* 读未提交
* 读提交
* 可重复读
* 串行化

---
