[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Network](https://mengxianbin.github.io/cs-notes/site/Network) /
[Protocol](https://mengxianbin.github.io/cs-notes/site/Network/Protocol) /
[HTTP](https://mengxianbin.github.io/cs-notes/site/Network/Protocol/HTTP) /
[应用](https://mengxianbin.github.io/cs-notes/site/Network/Protocol/HTTP/%E5%BA%94%E7%94%A8) /
[Cookie](https://mengxianbin.github.io/cs-notes/site/Network/Protocol/HTTP/%E5%BA%94%E7%94%A8/Cookie) /
[HttpOnly](https://mengxianbin.github.io/cs-notes/site/Network/Protocol/HTTP/%E5%BA%94%E7%94%A8/Cookie/HttpOnly)

* HttpOnly 的 Cookie 不能被 JavaScript 脚本调用

* 跨站脚本攻击（XSS）
    * 常常使用 JavaScript 的 `document.cookie` 窃取用户 Cookie
    * HttpOnly 可在一定程度上避免 XSS 攻击

```http
Set-Cookie: id=a3fWa; Expires=Wed, 21 Oct 2015 07:28:00 GMT; Secure; HttpOnly
```

---
