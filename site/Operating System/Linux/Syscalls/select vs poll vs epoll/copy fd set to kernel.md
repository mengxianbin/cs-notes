[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Operating System](https://mengxianbin.github.io/cs-notes/site/Operating%20System) /
[Linux](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux) /
[Syscalls](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Syscalls) /
[select vs poll vs epoll](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Syscalls/select%20vs%20poll%20vs%20epoll) /
[copy fd set to kernel](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Syscalls/select%20vs%20poll%20vs%20epoll/copy%20fd%20set%20to%20kernel)

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
