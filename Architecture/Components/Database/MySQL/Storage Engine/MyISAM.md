# MyISAM

* 设计简单
* 数据格式紧密
* 不支持事务
* 不支持行级锁，只能整表加锁
    * 读：共享锁
    * 写：拍他锁
    * 并发插入
        * 读操作时可以插入新记录

## 特性

* 压缩表
* 空间数据索引

## 使用场景

* 只读
* 表小

## 检查和修复

* 支持手动、自动
* 慢
* 可能丢数据

---

* DELAY_KEY_WRITE 选项
    * 不直接写磁盘，写缓冲区

---
