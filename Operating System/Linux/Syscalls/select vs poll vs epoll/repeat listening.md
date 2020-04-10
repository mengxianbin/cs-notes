# 重复监听

* select
    * copy new fd_set to kernel
* poll
    * copy new struct pollfd to kernel
* epoll
    * reuse old RB-Tree

---
