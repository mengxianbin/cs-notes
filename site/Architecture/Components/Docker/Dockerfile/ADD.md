[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Docker](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Docker) /
[Dockerfile](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Docker/Dockerfile) /
[ADD](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Docker/Dockerfile/ADD)

* 复制本地文件到容器
    * 类似 COPY
    * 额外副作用
        * 本地 tar 提取
        * 远程 URL 支持
    * 推荐用 ADD