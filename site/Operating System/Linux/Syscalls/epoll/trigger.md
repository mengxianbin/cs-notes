[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Operating System](https://mengxianbin.github.io/cs-notes/site/Operating%20System) /
[Linux](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux) /
[Syscalls](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Syscalls) /
[epoll](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Syscalls/epoll) /
[trigger](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Syscalls/epoll/trigger)

# Epoll 的触发方式

* edge trigger - ET
* level trigger - LT

* epoll_wait
    * 阻塞等待 IO 事件

* select, poll 只支持 LT 工作模式
* epoll 的默认工作模式为 LT 模式

## 水平触发的时机

* 缓冲区不空，LT 模式返回读就绪
* 缓冲区不满，LT 模式返回写就绪

## 边缘触发的时机

* 读
    * `空 -> 不空`
    * 新数据达到 （缓冲区内数据变多）
    * 缓冲区有数据可读，且进程对相应 fd 进行 epoll_ctl_mod 修改 epoll_in 事件时
* 写
    * `满 -> 不满`
    * 旧数据被发走 （缓冲区变少）
    * 缓冲区有空间可写，且进程对相应 fd 进行 epoll_ctl_mod 修改 epoll_out 事件时

---
