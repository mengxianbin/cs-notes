* 申请外部资源
    * 系统调用
    * 中断
    * 异常

## 例子

* 读写文件
    * open
    * read
        * 地址映射
            * 虚拟到物理
            * 缺页中断
    * write

* 申请内存
    * malloc
        * brk
        * mmap

---

## 系统调用类型

* 进程控制
    * fork
    * exit
* 文件管理
    * chmod
    * chown
    * chgrp
* 设备
    * open
    * read
    * write
* 信息（系统、硬件）
    * getcpu
* 通信
    * pipe
    * mmap

---
