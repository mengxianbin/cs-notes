[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Operating System](https://mengxianbin.github.io/cs-notes/site/Operating%20System) /
[Linux](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux) /
[Syscalls](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Syscalls) /
[select vs poll vs epoll](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Syscalls/select%20vs%20poll%20vs%20epoll) /
[intro](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Syscalls/select%20vs%20poll%20vs%20epoll/intro)

* select O(n)
* poll O(n)
* epoll O(1)

---

* epoll
    * event poll
    * 事件驱动
        * 每个事件关联 fd

---

* select, poll, epoll, 本质都是同步 IO
    * 需要在读写事件就绪后，自己负责进行读写
    * 读写过程阻塞

---

* select
    * 把 fd 从用户态拷贝到内核态，fd 多时，开销大
    * 每次调用，需要遍历所有 fd，开销大
    * 默认支持 1024 个 fd，太少了

* poll
    * 与 select 类似
    * 没有 fd 数量限制

* epoll
    * epoll_ctl
    * epoll_wait
        * 需要：轮询判空
        * 不需要：不需要遍历 fd
    * 使用 mmap 避免内核态拷贝

---
