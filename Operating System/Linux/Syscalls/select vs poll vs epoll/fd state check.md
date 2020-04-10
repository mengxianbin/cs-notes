## fd 状态监测

* select
    * 轮询所有 fd
    * return ready mask
    * 根据 mask 给 fd_set 赋值
* poll
    * 轮询 fd
        * 就绪的 fd 加入等待队列
* epoll
    * 回调
        * epoll_ctl add
            * add fd -> RB Tree
            * register callback
            * 内核监测到某 fd 可读写
                * 调用回调
                    * add fd -> ready fd list

---
