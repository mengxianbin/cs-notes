
```c
int epoll_create(int size)
```

* 内核产生一个 epoll 实例
* 返回一个文件描述符
    * epoll实例的句柄
        * epfd

---

```c
int epoll_ctl(int epfd， int op， int fd， struct epoll_event *event)

typedef union epoll_data {
    void *ptr; /* 指向用户自定义数据 */
    int fd; /* 注册的文件描述符 */
    uint32_t u32; /* 32-bit integer */
    uint64_t u64; /* 64-bit integer */
} epoll_data_t;

struct epoll_event {
    uint32_t events; /* 描述epoll事件 */
    epoll_data_t data; /* 见上面的结构体 */
};
```

* op
    * EPOLL_CTL_ADD：向interest list添加一个需要监视的描述符
    * EPOLL_CTL_DEL：从interest list中删除一个描述符
    * EPOLL_CTL_MOD：修改interest list中一个描述符

* epoll_event
    * events
        * bit mask
        * 期望事件
        * 可多选

* 常用 epoll 事件
    * EPOLLIN：fd 可读
    * EPOLLOUT：fd 可写
    * EPOLLET：edge triggered
    * EPOLLONESHOT：只首次通知，后续不通知
    * EPOLLHUP：fd 挂起（默认监测事件）
    * EPOLLERR：fd 错误 （默认检测事件）

---

```c
int epoll_wait(int epfd， struct epoll_event *events， int maxevents， int timeout)
```

---
