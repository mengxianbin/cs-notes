[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Operating System](https://mengxianbin.github.io/cs-notes/site/Operating%20System) /
[Linux](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux) /
[Syscalls](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Syscalls) /
[poll](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Syscalls/poll) /
[pollfd](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Syscalls/poll/pollfd)

```c
struct pollfd {
    int   fd;         /* file descriptor */
    short events;     /* requested events */
    short revents;    /* returned events */
};
```
