* 独立的列
    * 不能是复合表达式或参数 （不可变）

* 多列索引
    * 用于多条件索引

* 索引列顺序
    * 让选择性最强的索引列在前面

* 前缀索引
    * 用于
        * BLOB
        * TEXT
        * VARCHAR
    * 只索引开始的部分字符
    * 长度取决于索引选择性

* 覆盖索引
    * 索引包含所要查询的字段
    * 对于索引包含的字段
        * 可以只读索引，不读记录
            * 减少数据读取
            * 避免系统调用（耗时）
    * 对于 InnoDB 引擎，若辅助索引能覆盖查询，则不需要访问主索引

---
