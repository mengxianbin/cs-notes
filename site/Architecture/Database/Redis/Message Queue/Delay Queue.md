[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Database](https://mengxianbin.github.io/cs-notes/site/Architecture/Database) /
[Redis](https://mengxianbin.github.io/cs-notes/site/Architecture/Database/Redis) /
[Message Queue](https://mengxianbin.github.io/cs-notes/site/Architecture/Database/Redis/Message%20Queue) /
[Delay Queue](https://mengxianbin.github.io/cs-notes/site/Architecture/Database/Redis/Message%20Queue/Delay%20Queue)

* sorted set
    * timestamp as score
    * message as value
    * zadd key score value

* zrangebyscore
    * zrangebyscore key (now-delay) now
    * loop

---
