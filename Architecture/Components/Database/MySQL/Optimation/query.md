## 切分大查询

```sql
DELETE FROM messages WHERE create < DATE_SUB(NOW(), INTERVAL 3 MONTH);
```

```sql
rows_affected = 0
do {
    rows_affected = do_query(
    "DELETE FROM messages WHERE create  < DATE_SUB(NOW(), INTERVAL 3 MONTH) LIMIT 10000")
} while rows_affected > 0
```

## 分解大连接查询

* 将联合查询分解为每个表单独查询，在应用层中进行关联
    * 缓存
        * 单表缓存命中率更高
        * 单表缓存利用率更高
    * 减少锁竞争
    * 在应用层做连接，更容易对数据库拆分，更容易做到高性能和可伸缩
    * 查询本身效率提升

```sql
SELECT * FROM tag
JOIN tag_post ON tag_post.tag_id=tag.id
JOIN post ON tag_post.post_id=post.id
WHERE tag.tag='mysql';
```

```sql
SELECT * FROM tag WHERE tag='mysql';
SELECT * FROM tag_post WHERE tag_id=1234;
SELECT * FROM post WHERE post.id IN (123,456,567,9098,8904);
```

---
