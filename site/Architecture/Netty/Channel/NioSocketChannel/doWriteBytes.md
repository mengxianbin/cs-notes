[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[Channel](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Channel) /
[NioSocketChannel](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Channel/NioSocketChannel) /
[doWriteBytes](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Channel/NioSocketChannel/doWriteBytes)

```java
    @Override
    protected int doWriteBytes(ByteBuf buf) throws Exception {
        final int expectedWrittenBytes = buf.readableBytes();
        return buf.readBytes(javaChannel(), expectedWrittenBytes);
    }
```
