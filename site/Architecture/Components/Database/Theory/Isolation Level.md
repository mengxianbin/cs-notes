[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Database](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Database) /
[Theory](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Database/Theory) /
[Isolation Level](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Database/Theory/Isolation%20Level)

| Level \ Phenomena | Lost Updates | Dirty Reads | Non-repeatable Reads | Phantoms |
|-------------------|:------------:|:-----------:|:--------------------:|:--------:|
| Read Uncommited   |      O       |      X      |          X           |    X     |
| Read commited     |      O       |      O      |          X           |    X     |
| Repeatable Reads  |      O       |      O      |          O           |    X     |
| Serializable      |      O       |      O      |          O           |    O     |

---
