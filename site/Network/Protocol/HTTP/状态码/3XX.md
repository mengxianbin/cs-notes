[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Network](https://mengxianbin.github.io/cs-notes/site/Network) /
[Protocol](https://mengxianbin.github.io/cs-notes/site/Network/Protocol) /
[HTTP](https://mengxianbin.github.io/cs-notes/site/Network/Protocol/HTTP) /
[状态码](https://mengxianbin.github.io/cs-notes/site/Network/Protocol/HTTP/%E7%8A%B6%E6%80%81%E7%A0%81) /
[3XX](https://mengxianbin.github.io/cs-notes/site/Network/Protocol/HTTP/%E7%8A%B6%E6%80%81%E7%A0%81/3XX)

## 3XX 重定向

* 301 Moved Permanently
    * 永久性重定向

* 302 Found
    * 临时性重定向

* 303 See Other
    * 临时重定向（与 302 类似）
    * 明确要求客户端应采用 GET 方法获取资源
        * HTTP 协议规定 301、302 不允许把 POST 改成 GET
        * 大多数浏览器在 301、302、303 状态下都会把 POST 改成 GET

* 304 Not Modified
    * 如果请求头包含一些条件，如
        * If-Match
        * If-None-Match
        * If-Modified-Since
        * If-Unmodified-Since
        * If-Range
    * 如果不满足条件，就回返回 304 状态码

* 307 Temporary Redirect
    * 临时重定向（与 302 类似）
    * 浏览器不会要求把 POST 改成 GET

---
