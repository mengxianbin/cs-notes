[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[ByteBuf](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/ByteBuf) /
[summary](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/ByteBuf/summary) /
[vs ByteBuffer](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/ByteBuf/summary/vs%20ByteBuffer)

## 扩容

* ByteBuffer: 定长
* ByteBuf: 自动扩容
    * ensureWritable

---

## 位置指针

* ByteBuffer
    * (0, position) -[flip]-> (position, limit)
* ByteBuf
    * readable: readerIndex -> writerIndex
    * writable: writerIndex -> capacity()

---
