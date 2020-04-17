[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Database](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Database) /
[MySQL](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Database/MySQL) /
[Index](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Database/MySQL/Index) /
[Full Text Index](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Database/MySQL/Index/Full%20Text%20Index)

## 全文索引

* MyISAM 引擎支持
* 查找文本中的关键字，而不是判等
* 查找使用 Match Against，而不是 WHERE
* 倒排索引实现
    * 记录这关键词到所在文档的映射

* InnoDB 引擎 在 5.6.4 也开始支持全文索引

---
