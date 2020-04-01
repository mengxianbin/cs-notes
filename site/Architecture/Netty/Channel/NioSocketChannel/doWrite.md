[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[Channel](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Channel) /
[NioSocketChannel](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Channel/NioSocketChannel) /
[doWrite](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Channel/NioSocketChannel/doWrite)


```java
    @Override
    protected void doWrite(ChannelOutboundBuffer in) throws Exception {
        SocketChannel ch = javaChannel();
        int writeSpinCount = config().getWriteSpinCount();
        do {
            ...
            int nioBufferCnt = in.nioBufferCount();
            ...            switch (nioBufferCnt) {
                case 0:
                    // We have something else beside ByteBuffers to write so fallback to normal writes.
                    writeSpinCount -= doWrite0(in);
                    break;
                case 1: {
                    // Only one ByteBuf so use non-gathering write
                    // Zero length buffers are not added to nioBuffers by ChannelOutboundBuffer, so there is no need
                    // to check if the total size of all the buffers is non-zero.
                    ...
                    final int localWrittenBytes = ch.write(buffer);
                    ...
                    break;
                }
                default: {
                    // Zero length buffers are not added to nioBuffers by ChannelOutboundBuffer, so there is no need
                    // to check if the total size of all the buffers is non-zero.
                    // We limit the max amount to int above so cast is safe
                    ...
                    final long localWrittenBytes = ch.write(nioBuffers, 0, nioBufferCnt);
                    ...
                    break;
                }
            ...
        } while (writeSpinCount > 0);

        incompleteWrite(writeSpinCount < 0);        
    }
```
