[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Database](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Database) /
[MySQL](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Database/MySQL) /
[questions](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Database/MySQL/questions) /
[原子操作 - 实现 computeIfPresent 语义](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Database/MySQL/questions/%E5%8E%9F%E5%AD%90%E6%93%8D%E4%BD%9C%20-%20%E5%AE%9E%E7%8E%B0%20computeIfPresent%20%E8%AF%AD%E4%B9%89)

- k: private key / unique index

```sql
INSERT INTO test(k, v) VALUES ('a', 1) ON DUPLICATE KEY UPDATE v=v+2; 
```

```sql
REPLACE INTO test(k, v) VALUES ('d', 4);
```

---
