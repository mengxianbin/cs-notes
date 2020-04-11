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



---
