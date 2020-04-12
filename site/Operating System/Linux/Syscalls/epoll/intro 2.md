[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Operating System](https://mengxianbin.github.io/cs-notes/site/Operating%20System) /
[Linux](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux) /
[Syscalls](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Syscalls) /
[epoll](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Syscalls/epoll) /
[intro 2](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Syscalls/epoll/intro%202)

* share memory with kernel
    * avoid copying

* use event instead of loop

---

* epoll_ctl
    * 注册文件描述符
    * 文件描述符就绪时
        * 内核 callback 回调
        * 激活该 文件描述符
    * 调用 epoll_wait 时得到通知

---
