```c
int pselect(
    int nfds, 
    fd_set *readfds, 
    fd_set *writefds,
    fd_set *exceptfds, 
    const struct timespec *timeout,
    const sigset_t *sigmask);
```

```c
struct timespec {
        time_t   tv_sec;        /* seconds */
        long     tv_nsec;       /* nanoseconds */
};
```


## 与 select 区别

* timeout
    * select
        * timeval
            * seconds
            * milliseconds
        * 会为指示剩余时间而更新 timeout

    * pselect
        * timespec
            * seconds
            * nanoseconds
        * 不会更新 timeout

* sigmask
    * sigmask == null ?
        * pselect 行为同 select

---
