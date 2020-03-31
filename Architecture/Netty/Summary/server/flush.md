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
