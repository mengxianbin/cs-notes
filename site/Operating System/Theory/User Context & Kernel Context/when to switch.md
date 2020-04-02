[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Operating System](https://mengxianbin.github.io/cs-notes/site/Operating%20System) /
[Theory](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Theory) /
[User Context & Kernel Context](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Theory/User%20Context%20&%20Kernel%20Context) /
[when to switch](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Theory/User%20Context%20&%20Kernel%20Context/when%20to%20switch)

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
