## 缓存定义

* 本地副本

* 优点
    * 降低距离时延
    * 本地网络带宽通常大于远程带宽
    * 减少长途的数据传输，节省费用
    * 降低原始服务器符合

## 缓存命中/未命中

* 容量有限
    * 本地有副本：cache hit
    * 本地无副本：cache miss

---

* 缓存再验证
    * 验证副本与原始内容是否一致
        * 新鲜度检测
        * revalidation

    * If-Modified-Since (IMS)

    * 通常在用户发起请求，且副本足够旧时，才发起对副本的再验证

---

* 再验证命中
    * revalidate hit
    * 内容无变化
    * HTTP 304: Not Modified
    * 标记副本为新鲜
    * 比缓存命中慢
        * 需要与原始服务器核对
    * 比再验证未命中快
        * 不需要从原始服务器拉取对象
* 再验证未命中
    * 传输完整对象内容
    * HTTP 200: OK
* 删除
    * 原始服务器已将对象删除
    * HTTP 404: Not Found
    * 缓存将副本删除

---

* 命中率（cache hit rate）
    * 文档命中率（document hit rate）
        * 请求命中比率
    * 字节命中率
        * 对于大文件，访问次数少，流量占用高，统计字节命中率比文档命中率更有意义

---

* 区分命中/未命中
    * 响应码
        * 都是 200
        * 无法区分
    * 首部
        * Date
        * Age
        * 相当于判断数据版本

## 缓存拓扑结构

* 私有缓存
    * 单个用户
    * 浏览器内

* 公有缓存
    * 某用户群体共享
    * 缓存代理服务器（caching proxy server）

## 缓存处理步骤

1. 接收
    * 请求
1. 解析
    * URL
    * 首部
1. 查询
    * 本地有没有副本
1. 新鲜度检测
    * 如果数据足够旧，需要缓存再验证，判断是否需要更新
1. 创建响应
1. 发送
1. 日志

## 缓存相关首部

* Cache-Control
    * public: 响应可以被缓存（默认）
    * private: 浏览器可以缓存，不允许中继缓存
    * no-store: 禁止缓存
    * no-cache: 浏览器可以缓存，但是未经新鲜度检查之前不可使用
    * max-age: 相对时间，单位秒
        * Date + max-age 为绝对过期时间

* Expires
    * value: date
    * HTTP/1.0
    * 绝对时间

* Date
    * value: date
    * response 报文
    * 响应时的时间

* Last-Modified
    * value: date
    * response 报文
    * 最后被修改的时间

* If-Modified-Since
    * value: date
    * request 报文
    * 用于新鲜度检查
        * 如未修改，更新缓存文档的 Date，并提供缓存副本
        * 如已修改，重新缓存新的文档，并提供新缓存的副本

* If-None-Match
    * value: tag
    * request 报文
    * 告知服务器，这些版本有对应的缓存
    * 配合 Etag 使用
    * Response Code
        * 304: Not Modified
            * 响应头与 200 OK 的情况一样
        * 200: OK
            * Cache-Control
            * Content-Location
            * Date
            * ETag
            * Expires
            * Vary
        * 412: Precondition Failed
            * 能够引发服务器状态改变

* Pragma
    * value: no-cache 等
    * HTTP/1.0
    * 相当于 Cache-Control

---

### 过期时间相关响应首部

* Expires
    * HTTP/1.0+
    * 响应头
    * 绝对过期时间
    * 弊端：电脑本地时间可能不准确
    * 例
        * Expires: Thu, 15 Apr 2020 20:00:00 GMT

* Cache-Control
    * HTTP/1.1
    * 响应头
    * 相对过期时间
    * 优先级高于 Expires
    * max-age: 单位秒
    * 例
        * Cache-Control: max-age=135360000

* Pragma
    * HTTP/1.0
    * 很少用

---

### 再验证相关请求首部

> 常用

* If-Modified-Since
    * 配合 Last-Modified
    * 例
        * If-Modified-Since: Fri, 12 May 2006 18:53:33 GMT
* If-None-Match
    * 最后修改时间验证弊端
        * 周期性重写，内容不变
        * 已修改，但是不重要，不需要同步
        * 有些服务器无法判断最后修改日期
        * 文档修改时间周期低于 1s
    * 配合 Etag 使用
        * 如果已缓存标签和服务器文档的标签有所不同
            * If-None-Match 首部就会执行所请求的方法
    * 例
        * If-None-Match: "abcd1234...abcd1234" (Etag)
    * 如果 HTTP/1.1 服务器收到的请求同时有 If-Modified-Since 和 If-None-Match
        * 需要两个都满足才能返回 304

> 不常用

* If-Unmodified-Since
* If-Range
* If-Match

---
