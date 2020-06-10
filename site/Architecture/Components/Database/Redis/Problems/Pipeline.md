[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Database](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Database) /
[Redis](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Database/Redis) /
[Problems](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Database/Redis/Problems) /
[Pipeline](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Database/Redis/Problems/Pipeline)

* 独占连接
    * 该连接在批处理结束前，该连接的其他命令将阻塞等待
* 批处理
    * 多个命令，一次网络交互，提高性能
* 场景
    * 实时性要求不高

## 与 multi 比较

* pipeline
    * 客户端缓冲，与服务器一次交互
    * 不保证原子性
* multi
    * 服务器缓冲，每个命令与服务器交互一次
    * 保证原子性

---
