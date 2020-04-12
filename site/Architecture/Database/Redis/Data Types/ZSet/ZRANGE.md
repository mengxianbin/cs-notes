[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Database](https://mengxianbin.github.io/cs-notes/site/Architecture/Database) /
[Redis](https://mengxianbin.github.io/cs-notes/site/Architecture/Database/Redis) /
[Data Types](https://mengxianbin.github.io/cs-notes/site/Architecture/Database/Redis/Data%20Types) /
[ZSet](https://mengxianbin.github.io/cs-notes/site/Architecture/Database/Redis/Data%20Types/ZSet) /
[ZRANGE](https://mengxianbin.github.io/cs-notes/site/Architecture/Database/Redis/Data%20Types/ZSet/ZRANGE)

* start and stop are inclusive ranges
    * so for example ZRANGE myzset 0 1 will return both the first and the second element of the sorted set.

```redis
redis> ZADD myzset 1 "one"
(integer) 1
redis> ZADD myzset 2 "two"
(integer) 1
redis> ZADD myzset 3 "three"
(integer) 1
redis> ZRANGE myzset 0 -1
1) "one"
2) "two"
3) "three"
redis> ZRANGE myzset 2 3
1) "three"
redis> ZRANGE myzset -2 -1
1) "two"
2) "three"
redis> 
```

```redis
redis> ZRANGE myzset 0 1 WITHSCORES
1) "one"
2) "1"
3) "two"
4) "2"
redis> 
```
