## 问题

### 内存占用
   
* 内核 hash table
* 量级
    * 1万条TIME_WAIT的连接，消耗约 1M 内存
    * 可以忽略

### CPU 占用
   
* 每次分配新的随机端口需要遍历 bound ports
* 可以忽略

---

## 参数调优

* net.ipv4.tcp_timestamps
    * RFC 1323 
        * TCP Reliability
            * TCP option
                * timestamp
                    * last send (4 Byte)
                    * last receive (4 Byte)
                    * Used By
                        * tcp_tw_reuse
                        * tcp_tw_recycle

* net.ipv4.tcp_tw_reuse
    * reuse time_wait connection
    * 场景
        * 不断与服务器建立短连接
        * 总是自己先发起关闭，处于 TIME_WAIT
        * 不断重新连接对方
    * 依赖 tcp_timestamps
        * 连接复用后重置时间戳
        * 延迟数据到达时，时间戳小于连接重置时间戳，那么该延迟数据可以安全丢弃
    * 需要连接双方同时支持 timestamps
    * 仅影响 outbound 连接
        * 即客户端连向服务器的 socket

* net.ipv4.tcp_tw_recycle
    * timewait 回收
    * 开启此配置后，内核快速回收 timewait 连接
        * 不再是 2MSL
        * 而是 RTO
            * Retransmission Timeout 数据包重传超时
            * 根据 RTT 动态计算，远小于 2MSL
    * socket 进入 TIME_WAIT 状态后，内核记录 socket 统计数据
        * 包括最后一次收到数据包的时间戳
        * 数据包到达时，如果时间晚于所记录的时间戳，就丢弃
    * 该配置需要连接双方支持 timestamps
        * 主要影响 inbound 连接
        * 服务端主动关闭连接后，快速回收 TIME_WAIT 状态的连接
    * 如果客户端处于 NAT 网络
        * 多个客户端，同一个 IP 出口
        * 在一个 RTO 时间内，只有一个客户端能与服务器连接成功
            * 不同的客户端发包时间不一致，造成服务器丢弃数据包 

---
