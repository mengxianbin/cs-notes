[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Database](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Database) /
[Redis](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Database/Redis) /
[LOCK](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Database/Redis/LOCK) /
[commands](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Database/Redis/LOCK/commands)

```redis

* SET key val [NX|XX] [EX seconds | PX milliseconds]
    * 可选参数 NX|XX：
        * NX表示只在键不存在时，才对键进行操作
        * XX表示只在键存在时对键进行操作
        * 缺省方式是NX
    * 可选参数 EX|PX：
        * 键过期的时间单位，后面跟长整型数字表示过期时间
        * EX表示秒
        * PX表示毫秒
        * 缺省不设置过期时间

* SETNX key val 
    * key 存在
        * 设置 value
        * 返回 1
    * key 不存在
        * 不设置 value
        * 返回 0

* GETSET key val
    * 设置新值
    * 返回旧值

* INCR key
    * 自增

* DEL key
    * 删除

```
