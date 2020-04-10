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
