[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Message Queue](https://mengxianbin.github.io/cs-notes/site/Architecture/Message%20Queue) /
[Compare](https://mengxianbin.github.io/cs-notes/site/Architecture/Message%20Queue/Compare)

## 一、资料文档

* ActiveMQ  多
* RabbitMQ  多
* Kafka     中
* RocketMQ  少

## 二、开发语音

* ActiveMQ  Java
* RabbitMQ  Erlang
* Kafka     Scala
* RocketMQ  Java

## 三、支持的协议

* ActiveMQ  OpenWire、STOMP、REST、XMPP、AMQP
* RabbitMQ  AMQP
* Kafka     自己定义的一套...（基于TCP）   
* RocketMQ  自己定义的一套...

## 四、消息存储

* ActiveMQ  内存、磁盘、数据库。支持少量堆积。
* RabbitMQ  内存、磁盘。支持少量堆积。
* Kafka     内存、磁盘、数据库。支持大量堆积。
* RocketMQ  磁盘。支持大量堆积。

## 五、消息事务

* ActiveMQ  支持
* RabbitMQ  支持
* Kafka     支持
* RocketMQ  支持

## 六、负载均衡

* ActiveMQ  支持
* RabbitMQ  支持不好
* Kafka     支持
* RocketMQ  支持

## 七、集群方式

* ActiveMQ  支持简单集群，高级集群模式支持不好
* RabbitMQ  支持简单集群，高级集群模式支持不好
* Kafka     Leader-Slave, 每台服务器既是Master也是Slave。
* RocketMQ  常用 多对 'Master-Slave' 模式，开源版本需手动切换Slave变成Master

## 八、管理界面

* ActiveMQ  一般
* RabbitMQ  好
* Kafka     一般
* RocketMQ  无

## 九、可用性

·高（主从）
·高（主从）
·非常高（分布式）
·非常高（分布式）

## 十、消息重复

* ActiveMQ  支持at least once
* RabbitMQ  支持at least once、at most once
* Kafka     支持at least once、at most once
* RocketMQ  支持at least once

## 十一、吞吐量 TPS

* ActiveMQ  比较大
* RabbitMQ  比较大
* Kafka     极大
* RocketMQ  大

## 十二、订阅形式和消息分发

* ActiveMQ  点对点(p2p)、广播（发布-订阅）
* RabbitMQ  提供了4种：direct, topic ,Headers和fanout。
* Kafka     基于topic以及按照topic进行正则匹配的发布订阅模式。
* RocketMQ  基于topic/messageTag以及按照消息类型、属性进行正则匹配的发布订阅模式

## 十三、顺序消息

* ActiveMQ  不支持
* RabbitMQ  不支持
* Kafka     支持
* RocketMQ  支持

## 十四、消息确认

* ActiveMQ  支持
* RabbitMQ  支持
* Kafka     支持
* RocketMQ  支持

## 十五、消息回溯

* ActiveMQ  不支持
* RabbitMQ  不支持
* Kafka     支持指定分区offset位置的回溯。
* RocketMQ  支持指定时间点的回溯。

## 十六、消息重试

* ActiveMQ  不支持
* RabbitMQ  不支持，但是可以利用消息确认机制实现。
* Kafka     不支持，但是可以实现。
* RocketMQ  支持

## 十七、并发度

* ActiveMQ  高
* RabbitMQ  极高
* Kafka     高
* RocketMQ  高

---

## 总结

* 都支持的特性
    * 消息存储
    * 消息事务
    * 消息确认
    * 消息重复
    * 集群
    * 管理界面
    * 可用性
    * 高吞吐
    * 高并发
    * 负载均衡

* RocketMQ 独有或突出
    * 支持消息大量堆积
    * 消息可按标签订阅
    * 顺序消息
    * 消息回溯 - 时间点
    * 消息重试

* RocketMQ 开源阉割
    * 手动切换 master
    * 无界面

---
