[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Database](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Database) /
[MySQL](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Database/MySQL) /
[Replication](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Database/MySQL/Replication)

## 复制

### 主从复制

* binlog 线程
    * 将主服务器的数据变更写入二进制日志
* IO 线程
    * 从主服务器读取 binlog，写入从服务器的中继日志 relay log
* SQL 线程
    * 读取中继日志，解析数据变更，在从服务器中重放

### 读写分离

#### 分工

* 主服务器处理写操作、以及实时性要求高的读操作
* 从服务器处理读操作

#### 分析

* 主从读写分离，缓解锁争用
* 从服务器可以使用 MyISAM
    * 提升查询性能
    * 节省系统开销
* 增加冗余，提高可用性

#### 实现

* 代理
    * 查询请求发往代理服务器
    * 代理服务器决定转发给谁

---
