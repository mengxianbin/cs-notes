* 响应报文中，第一行为状态行
    * 包含状态码 和 原因短语
    * 用于告知客户端请求结果

---

| 状态码 | 类别                    | 含义                       |
|--------|-------------------------|----------------------------|
| 1XX    | Infomational 信息       | 接收的请求正在处理         |
| 2XX    | Success 成功            | 请求正常处理完毕           |
| 3XX    | Redirection 重定向      | 需要进行附加操作以完成请求 |
| 4XX    | Client Error 客户端错误 | 服务器无法处理请求         |
| 5XX    | Server Error 服务器错误 | 服务器处理请求出错         |

---
