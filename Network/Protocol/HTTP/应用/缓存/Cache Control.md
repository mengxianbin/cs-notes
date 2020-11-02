* HTTP 1.1
* Cache-Control 首部字段

---

## 禁止缓存

* no-store 指令
    * 不能缓存任一部分
        * 请求
        * 响应

> Cache-Control: no-store

## 强制确认缓存

* no-cache 指令
    * 验证资源有效性
        * 缓存服务器 -> 源服务器
        * 仅当缓存资源有效时
            * 才能使用该缓存对客户端的请求进行响应。

> Cache-Control: no-cache

## 私有缓存、共有缓存

* private 指令
    * 指定缓存资源为私有缓存
    * 单独用户使用
    * 一般存储在用户浏览器

> Cache-Control: private

* public 指令
    * 指定公有缓存
    * 多用户使用
    * 存于代理服务器

> Cache-Control: public

## 缓存过期机制

* max-age 指令
    * 请求报文
    * 响应报文
        * 缓存服务器

> Cache-Control: max-age=31536000

* Expires 首部字段
    * 告知缓存服务器
        * 该资源什么时候会过期

> Expires: Wed, 04 Jul 2012 08:26:05 GMT

* 在 HTTP/1.1 中，会优先处理 max-age 指令
* 在 HTTP/1.0 中，max-age 指令会被忽略掉

---
