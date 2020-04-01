[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Docker](https://mengxianbin.github.io/cs-notes/site/Architecture/Docker) /
[unionFS](https://mengxianbin.github.io/cs-notes/site/Architecture/Docker/unionFS)

* unionFS
    * 可以挂载多个物理隔离的目录到同一目录下
    * 不同内核版本，实现技术可能不同
    * bootfs
    * rootfs

* bootfs
    * boot file system
    * 启动完成后，整个 Linux 内核加载进内存，bootfs 卸载，释放内存
    * 同一内核版本，所有发行版 bootfs 相同

* rootfs
    * root file system
    * 同一内核版本，不同发行版 rootfs 可能不同
        * 配置文件
        * 可执行文件
        * 库文件
        * 目录结构
            * /bin
            * /boot
            * /dev
            * /etc
            * /home
            * /lib
            * /mnt
            * /opt
            * /proc
            * /root
            * /tmp
            * /usr
            * /var
    * 启动时挂载为只读
    * 启动完成后，修改为读写

* 容器分层结构
    * 内核
    * 基础镜像
    * 若干镜像层
    * 容器读写层

* 容器启动后，容器内到改动
    * 只作用于容器内到可读写层
    * 对下面对只读层无影响
