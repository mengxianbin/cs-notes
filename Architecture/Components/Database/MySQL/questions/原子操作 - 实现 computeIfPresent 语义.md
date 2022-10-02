- k: private key / unique index

```sql
INSERT INTO test(k, v) VALUES ('a', 1) ON DUPLICATE KEY UPDATE v=v+2; 
```

```sql
REPLACE INTO test(k, v) VALUES ('d', 4);
```

---
