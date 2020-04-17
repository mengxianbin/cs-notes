[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Distribution](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Distribution) /
[Lock](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Distribution/Lock) /
[Redis](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Distribution/Lock/Redis)

* setnx 指令
    * set if not exist
    * return true if not exist
        * set successfully

* expire 指令
    * 避免锁释放失败

---

## RedLock 算法

* 多个 Redis 实例，提高可用性
    * 尝试从 N 个相互独立的 Redis 实例获取锁
    * 计算锁获取的时间
        * 成功条件
            * 小于锁过期时间
            * 从过半实例获得了锁
                * N/2 + 1
    * 如果获取失败，到每个节点释放锁

---

## 重入锁

* 维护进程内的锁持有线程
* 如果加锁线程是当前线程，说明已重入
    * 降级为本地锁
        * 增加锁持有数

---
