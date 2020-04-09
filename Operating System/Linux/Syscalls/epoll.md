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
