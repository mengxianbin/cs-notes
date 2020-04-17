[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Database](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Database) /
[Redis](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Database/Redis) /
[Transaction](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Database/Redis/Transaction) /
[EXEC](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Database/Redis/Transaction/EXEC)

* 若干命令加入队列
* 收到 EXEC 命令，开始执行队列中的命令
    * 中间命令失败时
        * 前面的命令不回滚
        * 后面的命令仍然继续执行
* 事务执行过程中，其他命令不会插入到命令序列中
