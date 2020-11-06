[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Container](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Container) /
[Docker](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Container/Docker) /
[Theory](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Container/Docker/Theory) /
[namespace](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Container/Docker/Theory/namespace)

| ns type | syscall arg   | kernel version |
|---------|---------------|----------------|
| Mount   | CLONE_NEWNS   | 2.4.19         |
| UTS     | CLONE_NEWUTS  | 2.6.19         |
| IPC     | CLONE_NEWIPC  | 2.6.19         |
| PID     | CLONE_NEWPID  | 2.6.24         |
| Network | CLONE_NEWNET  | 2.6.29         |
| User    | CLONE_NEWUSER | 3.8            |

---
