[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Network](https://mengxianbin.github.io/cs-notes/site/Network) /
[commands](https://mengxianbin.github.io/cs-notes/site/Network/commands) /
[netstat](https://mengxianbin.github.io/cs-notes/site/Network/commands/netstat)

* netstat
    * show net status

* 查看 TCP 各个状态的数量


```sh
netstat -n | awk '/^tcp/ {++state[$NF]} END {for(key in state) print key,"\t",state[key]}'
```

```
TIME_WAIT    5856
CLOSE_WAIT   268
FIN_WAIT1    3
ESTABLISHED  4837
SYN_RECV     14
CLOSING      1
```

---

* 只看正常的并发连接
    * netstat -nat|grep ESTABLISHED|wc -l

---
