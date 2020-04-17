[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Summary](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Summary) /
[server](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Summary/server) /
[flush](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Summary/server/flush)

* flush - Channel.Unsafe
    * addFlush - ChannelOutboundBuffer
    * flush0 - AbstractChannel.AbstractUnsafe
        * doWrite - AbstractChannel / NioSocketChannel
            * doWrite0 - AbstractNioByteChannel
                * doWriteInternal - AbstractNioByteChannel
                    * doWriteBytes - NioSocketChannel
                        * javaChannel - NioSocketChannel
                        * readBytes - ByteBuf
                            * getBytes - ByteBuf / UnpooledDirectByteBuf
                                * write - GatheringByteChannel / SocketChannel(java)

---
