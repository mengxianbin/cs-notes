| Level \ Phenomena | Lost Updates | Dirty Reads | Non-repeatable Reads | Phantoms |
|-------------------|:------------:|:-----------:|:--------------------:|:--------:|
| Read Uncommited   |      O       |      X      |          X           |    X     |
| Read commited     |      O       |      O      |          X           |    X     |
| Repeatable Reads  |      O       |      O      |          O           |    X     |
| Serializable      |      O       |      O      |          O           |    O     |

---
