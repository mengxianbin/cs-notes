[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[ByteBuf](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/ByteBuf) /
[summary](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/ByteBuf/summary) /
[Hierarchy](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/ByteBuf/summary/Hierarchy)

* Object
    * ByteBuf
        * AbstractByteBuf

            * AbstractDerivedByteBuf
                * AbstractUnpooledSlicedByteBuf
                    * UnpooledSlicedByteBuf
                    * ~~SlicedByteBuf~~ @See `ByteBuf.slice(int index, int length)`

            * AbstractReferenceCountedByteBuf             

                * CompositeByteBuf

                * AbstractPooledDerivedByteBuf
                    * PooledSlicedByteBuf

                * PooledByteBuf
                    * PooledDirectByteBuf
                    * PooledUnsafeDirectByteBuf
                    * PooledHeapByteBuf
                        * PooledUnsafeHeapByteBuf
                * UnpooledDirectByteBuf
                    * UnpooledUnsafeDirectByteBuf
                * UnpooledHeapByteBuf
                    * UnpooledUnsafeHeapByteBuf

---

## 使用场景

### 分配

* Heap
    * 内存的分配和回收都在堆，速度很快
    * Socket的IO读写，需要把堆内存对应的缓冲区复制到内核Channel中，内存复制会影响性能
* Direct
    * 内存的分配在非堆（方法区），不需要内存复制，IO读取的速度快
    * 内存的分配较慢

### 回收

* Pooled
    * 循环利用创建的 ByteBuf，提高内存使用效率
* Unpooled
    * 负载不高时，可以考虑使用

## 零拷贝

* CompositeByteBuf
* ByteBuf.slice(int index, int length)
* FileChannel.transferTo
    * java.nio.channels

---
