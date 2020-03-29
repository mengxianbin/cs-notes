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
