[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Database](https://mengxianbin.github.io/cs-notes/site/Architecture/Database) /
[Redis](https://mengxianbin.github.io/cs-notes/site/Architecture/Database/Redis) /
[Cluster](https://mengxianbin.github.io/cs-notes/site/Architecture/Database/Redis/Cluster) /
[Hash](https://mengxianbin.github.io/cs-notes/site/Architecture/Database/Redis/Cluster/Hash)

* 一致性哈希
    * 哈希偏斜
        * 虚拟节点：将每个物理节点映射为多个虚拟节点

* 适用于内存缓存
    * 因为失联的持久化节点无法自动在哈希环上迁移

---

* 哈希槽
    * 划分数据区间
    * 每个区间有一组主从实例对应
        * 一主多从

* 适用于持久化

---

* 一致性哈希和哈希槽可以结合使用

---

## 一致性哈希

* 示例
    * ip: port # virtualN

* 哈希算法优劣标准
    * 平衡性 - Balance
        * 哈希结果尽可能分布到所有缓冲中去
            * 提高空间利用率
    * 单调性 - Monotonicity
        * 同一内容可以分到原分区，或新分区，而不是其他旧分区
    * 分散性 - Spread
        * 避免不同的终端将同样的内容映射到不同的分区
    * 负载 - Load
        * 同一缓冲区，被不同用户映射为不同内容

---
