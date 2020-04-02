[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Redis](https://mengxianbin.github.io/cs-notes/site/Architecture/Redis) /
[Persistance](https://mengxianbin.github.io/cs-notes/site/Architecture/Redis/Persistance) /
[AOF](https://mengxianbin.github.io/cs-notes/site/Architecture/Redis/Persistance/AOF) /
[strategy](https://mengxianbin.github.io/cs-notes/site/Architecture/Redis/Persistance/AOF/strategy)

* always
* everysec
* no

---

| command  | pro            | con          |
|----------|----------------|--------------|
| always   | 不丢数据       | IO 开销大    |
| everysec | 最多丢 1s 数据 | 会丢 1s 数据 |
| no       | 不用管         | 不可控       |

---
