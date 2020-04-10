## fd 拷贝方式

* select
    * fd_set * 3: 读、写、异常
    * 单进程 fd 数量限制：默认 1024
* poll
    * struct pollfd: 用户 -> 内核
* epoll
    * epoll_create
        * 在内核高速 cache 区
            * create: RB Tree * 1
            * create: ready fd list * 1
    * epoll_ctl
        * 添加 fd 会在 RB Tree 增加相应节点

---
