[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Operating System](https://mengxianbin.github.io/cs-notes/site/Operating%20System) /
[Linux](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux) /
[Syscalls](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Syscalls) /
[select vs poll vs epoll](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Syscalls/select%20vs%20poll%20vs%20epoll) /
[repeat listening](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Syscalls/select%20vs%20poll%20vs%20epoll/repeat%20listening)

# 重复监听

* select
    * copy new fd_set to kernel
* poll
    * copy new struct pollfd to kernel
* epoll
    * reuse old RB-Tree

---
