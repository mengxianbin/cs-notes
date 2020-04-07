[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Network](https://mengxianbin.github.io/cs-notes/site/Network) /
[HTTP](https://mengxianbin.github.io/cs-notes/site/Network/HTTP) /
[应用](https://mengxianbin.github.io/cs-notes/site/Network/HTTP/%E5%BA%94%E7%94%A8) /
[Cookie](https://mengxianbin.github.io/cs-notes/site/Network/HTTP/%E5%BA%94%E7%94%A8/Cookie) /
[Session](https://mengxianbin.github.io/cs-notes/site/Network/HTTP/%E5%BA%94%E7%94%A8/Cookie/Session)

* 用户信息
    * 可以存于客户端浏览器 Cookie
    * 也可以存于服务器 Session
        * 更安全

* Session 存储形式
    * 内存
    * 文件
    * 数据库
    * Redis

* Session 维护登录状态的过程
    * 用户登录
        * 提交包含用户名、密码的表单，放入 HTTP 请求报文中
    * 服务器验证该用户的用户名、密码
        * 如果正确，用户信息存于 Redis，Key 为 Session ID
    * 服务器返回响应报文
        * Set-Cookie 头字段包含 该 Session ID
    * 客户端收到响应报文
        * 将 Cookie 存于浏览器
    * 客户端再次请求同一服务器时
        * 从浏览器取出相应 Cookie，随请求发往服务器
    * 服务器由 Redis 取出相应 Session ID 的用户信息

---

* Session ID 安全性
    * 不能轻易获取，避免恶意攻击
    * ID 值不能过于简单
    * 经常重新生成
    * 极敏感场景要求重新验证
        * 密码
        * 验证码

---
