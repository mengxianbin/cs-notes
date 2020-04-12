[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Operating System](https://mengxianbin.github.io/cs-notes/site/Operating%20System) /
[Linux](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux) /
[Syscalls](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Syscalls) /
[select vs poll vs epoll](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Syscalls/select%20vs%20poll%20vs%20epoll) /
[ready fd set to user](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Syscalls/select%20vs%20poll%20vs%20epoll/ready%20fd%20set%20to%20user)

# 就绪 fd 如何回传用户态

* select
    * copy fd_set to user context
    * return fd count
    * user context dont know which fd is ready, have to traverse
* poll
    * copy fd array to user context
    * return fd count
    * user context dont know which fd is ready, have to traverse
* epoll
    * check whether the ready fd list is empty
    * memory share
        * between user and kernel
        * avoid copying

---
