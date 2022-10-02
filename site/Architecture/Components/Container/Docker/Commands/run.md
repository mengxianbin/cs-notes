[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Container](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Container) /
[Docker](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Container/Docker) /
[Commands](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Container/Docker/Commands) /
[run](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Container/Docker/Commands/run)

* --restart
    * no
    * failure
    * always
        * 自动重启，即便 stopped
        * 手动 stop, 不再重启
        * 重启 Docker Daemon, container 自动重启
    * unless-stopped
        * 类似 always
        * 一旦手动 stop, 重启 daemon 也不会使 container 重启

---
