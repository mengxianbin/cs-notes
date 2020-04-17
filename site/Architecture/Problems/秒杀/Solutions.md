[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Problems](https://mengxianbin.github.io/cs-notes/site/Architecture/Problems) /
[秒杀](https://mengxianbin.github.io/cs-notes/site/Architecture/Problems/%E7%A7%92%E6%9D%80) /
[Solutions](https://mengxianbin.github.io/cs-notes/site/Architecture/Problems/%E7%A7%92%E6%9D%80/Solutions)

## 服务单一职责

* 服务拆分（微服务）

---

## 数据库

* 分库分表
    * 秒杀库与其他库分离
    * 冷热分离

---

## 秒杀链接加盐

* URL 
    1. 拼接动态串
    1. 加密

---

## 库存预热

* Redis 缓存

---

## Redis 集群

* 开启持久化
* 主从同步
* 读写分离
* 哨兵

## Redis Lua

> @Since Redis v2.6
> CAS 高效执行
> 可以组合多个命令

* Lua 脚本类似 Redis 事务，具备原子性

---

## Nginx

> Nginx: 高性能 Web 服务器，单台支持 `万` 级并发
> Tomcat: 支持 `百` 级并发

* 服务节点大量部署
* 负载均衡

---

## 网关

* 拦截恶意请求
    * 同一 用户 高频请求

---

## 资源静态化

* CDN

---

## 限流

* 前端限流
    * 秒杀开始前置灰
    * 一次请求后，按钮置灰等待
* 后端限流
    * 功能
        * 秒杀一旦结束
            * 关闭开关变量，将请求拦在最外层
            * 通知客户端按钮置灰
    * 组件
        * Sentinel
        * Hystrix

---

## 消息队列

* 削峰填谷
* 异步

---

## 止损

* 限流
* 熔断
* 降级
* 隔离

---
