[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Operating System](https://mengxianbin.github.io/cs-notes/site/Operating%20System) /
[Linux](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux) /
[Syscalls](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Syscalls) /
[select](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Syscalls/select) /
[select](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Syscalls/select/select)

```c
int select(int nfds,
    fd_set *readfds,
    fd_set *writefds,
    fd_set *exceptfds,
    struct timeval *timeout);
```

```c
struct timeval {
    time_t         tv_sec;     /* seconds */
    suseconds_t    tv_usec;    /* microseconds */
};
```

