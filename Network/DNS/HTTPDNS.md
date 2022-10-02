* LocalDNS 的问题
    * 运营商劫持
        * 插入广告
        * 增加时延
        * 解析结果不按 TTL 缓存
        * 解析被错误递归（跨地区甚至跨运营商）

* 解决
    * 绕开运营商来做 domain-IP 映射
        * HTTPDNS

---

* HTTPDNS
    * 使用 IP 直接发 HTTP 请求 目标服务器的A记录地址
        * 避免运营商劫持
        * 省去 domain 到 IP 的解析
            * 减少延迟

---
