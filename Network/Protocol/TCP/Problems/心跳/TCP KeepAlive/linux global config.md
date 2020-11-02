* /etc/sysctl.conf

```conf
net.ipv4.tcp_keepalive_intvl = 20 # 探测间隔（秒）
net.ipv4.tcp_keepalive_probes = 3 # 探测尝试次数
net.ipv4.tcp_keepalive_time = 60  # Idle （秒）
```
