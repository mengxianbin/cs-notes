[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Network](https://mengxianbin.github.io/cs-notes/site/Network) /
[Protocol](https://mengxianbin.github.io/cs-notes/site/Network/Protocol) /
[TCP](https://mengxianbin.github.io/cs-notes/site/Network/Protocol/TCP) /
[Problems](https://mengxianbin.github.io/cs-notes/site/Network/Protocol/TCP/Problems) /
[心跳](https://mengxianbin.github.io/cs-notes/site/Network/Protocol/TCP/Problems/%E5%BF%83%E8%B7%B3) /
[TCP KeepAlive](https://mengxianbin.github.io/cs-notes/site/Network/Protocol/TCP/Problems/%E5%BF%83%E8%B7%B3/TCP%20KeepAlive) /
[linux global config](https://mengxianbin.github.io/cs-notes/site/Network/Protocol/TCP/Problems/%E5%BF%83%E8%B7%B3/TCP%20KeepAlive/linux%20global%20config)

* /etc/sysctl.conf

```conf
net.ipv4.tcp_keepalive_intvl = 20 # 探测间隔（秒）
net.ipv4.tcp_keepalive_probes = 3 # 探测尝试次数
net.ipv4.tcp_keepalive_time = 60  # Idle （秒）
```
