[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Redis](https://mengxianbin.github.io/cs-notes/site/Architecture/Redis) /
[features](https://mengxianbin.github.io/cs-notes/site/Architecture/Redis/features)

## Pros

* 数据类型丰富
* 支持集群
* 支持持久化
* 效率高

## Cons

* 单进程单线程，长命令会导致 Redis 阻塞
* 集群下多 key 操作（事务、MGET、MSET）无法使用 无法自动迁移

---
