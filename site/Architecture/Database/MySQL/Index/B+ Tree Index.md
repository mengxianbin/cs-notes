[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Database](https://mengxianbin.github.io/cs-notes/site/Architecture/Database) /
[MySQL](https://mengxianbin.github.io/cs-notes/site/Architecture/Database/MySQL) /
[Index](https://mengxianbin.github.io/cs-notes/site/Architecture/Database/MySQL/Index) /
[B+ Tree Index](https://mengxianbin.github.io/cs-notes/site/Architecture/Database/MySQL/Index/B+%20Tree%20Index)

* Balance Tree

* Operation
    * Deep Search
        * break balance
        * tree node split
        * tree node re-merge

---

## Compare with Red-Black Tree

* less search
    * bigger out degree
        * more children in one node
    * smaller tree height

* use disk pre-read `ordered reading`
    * less IO operation
    * less disk track switching

---

## Primary Index

* unique each table
* save data in `data` field

## Secondary Index

* save primary key in `data` field

---
