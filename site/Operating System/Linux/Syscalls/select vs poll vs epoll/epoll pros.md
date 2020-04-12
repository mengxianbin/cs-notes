[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Operating System](https://mengxianbin.github.io/cs-notes/site/Operating%20System) /
[Linux](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux) /
[Syscalls](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Syscalls) /
[select vs poll vs epoll](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Syscalls/select%20vs%20poll%20vs%20epoll) /
[epoll pros](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Syscalls/select%20vs%20poll%20vs%20epoll/epoll%20pros)

# epoll 更高效的原因

1. fd 数据结构
    * select
        * fd bit marks: fd count limit - 1024
    * poll
        * pollfd[]: custom-length array

1. which fds are ready?
    * select, poll dont know
    * epoll know

1. fd copy
    * select, poll copy between user and kernel
    * epoll use shared memory between user and kernel

1. trigger
    * select, poll: level trigger
    * epoll: edge trigger

---
