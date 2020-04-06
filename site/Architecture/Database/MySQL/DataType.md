[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Database](https://mengxianbin.github.io/cs-notes/site/Architecture/Database) /
[MySQL](https://mengxianbin.github.io/cs-notes/site/Architecture/Database/MySQL) /
[DataType](https://mengxianbin.github.io/cs-notes/site/Architecture/Database/MySQL/DataType)

## 整形

| type      | bits |
|-----------|------|
| TINYINT   | 8    |
| SMALLINT  | 16   |
| MEDIUMINT | 24   |
| INT       | 32   |
| BIGINT    | 64   |

* INT(11)
    * 规定交互工具显示字符个数
    * 对于存储没有意义

## 浮点数

* FLOAT
* DOUBLE

* DECIMAL
    * 10进制
    * 比 Float 耗时

* FLOAT, DOUBLE, DECIMAL 都支持指定列宽
    * DECIMAL(18, 9)
        * 9 位 小数
        * 18 - 9 位 整数

## 字符串

* CHAR: 定长
* VARCHAR: 变长
    * 只存必要内容，节省空间
    * update 需要分页时会耗时
        * MyISAM： 拆分片段
        * InnoDB： 分裂页

## 时间、日期

* DATETIME
    * 8字节
    * 精度秒
    * 时区无关
    * 范围：1000 ~ 9999
    * 显示：ANSI 标准
        * “2008-01-16 22:37:08”

* TIMESTAMP
    * UNIX 时间戳（秒）
        * 格林威治
        * 起始：1970-01-01 00:00:00
    * 4 字节
    * 只能表示 1970 ~ 2038
    * 时区相关
        * 一个时间戳在不同时区代表不同时间
    * 转换
        * FROM_UNIXTIME()
        * UNIX_TIMESTAMP()
    * 默认值为当前时间

* 尽量使用 TIMESTAMP
    * 比 DATETIME 空间效率更高

---
