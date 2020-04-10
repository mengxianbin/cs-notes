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
