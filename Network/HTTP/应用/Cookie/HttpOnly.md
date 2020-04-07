* HttpOnly 的 Cookie 不能被 JavaScript 脚本调用

* 跨站脚本攻击（XSS）
    * 常常使用 JavaScript 的 `document.cookie` 窃取用户 Cookie
    * HttpOnly 可在一定程度上避免 XSS 攻击

```http
Set-Cookie: id=a3fWa; Expires=Wed, 21 Oct 2015 07:28:00 GMT; Secure; HttpOnly
```

---
