[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Database](https://mengxianbin.github.io/cs-notes/site/Architecture/Database) /
[Redis](https://mengxianbin.github.io/cs-notes/site/Architecture/Database/Redis) /
[Persistance](https://mengxianbin.github.io/cs-notes/site/Architecture/Database/Redis/Persistance) /
[RDB](https://mengxianbin.github.io/cs-notes/site/Architecture/Database/Redis/Persistance/RDB) /
[command](https://mengxianbin.github.io/cs-notes/site/Architecture/Database/Redis/Persistance/RDB/command)

| command    | save         | bgsave                |
|------------|--------------|-----------------------|
| IO type    | sync         | async                 |
| blocking   | yes          | yes (when fork)       |
| complexity | O(n)         | O(n)                  |
| pro        | save memory  | not block client      |
| con        | block client | cost memory when fork |

---
