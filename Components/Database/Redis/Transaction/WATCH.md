* 监视一个(或多个) key
* 如果在事务执行之前这个(或这些) key 被其他命令所改动，那么事务将被打断

---

* WATCH 可以用于创建 Redis 没有内置的原子操作

* zpop
    * 原子弹出有序集合中，分数最小的元素

    ```redis
    WATCH zset
    element = ZRANGE zset 0 0
    MULTI
        ZREM zset element
    EXEC
    ```

    * 程序需要重复执行这段代码
        * 直到 EXEC 的返回值不是`空多条回复（null multi-bulk reply）`即可

---
