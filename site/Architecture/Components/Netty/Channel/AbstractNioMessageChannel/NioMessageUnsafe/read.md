[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Channel](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel) /
[AbstractNioMessageChannel](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel/AbstractNioMessageChannel) /
[NioMessageUnsafe](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel/AbstractNioMessageChannel/NioMessageUnsafe) /
[read](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel/AbstractNioMessageChannel/NioMessageUnsafe/read)

* read
    * final ChannelPipeline pipeline = pipeline();
    * pipeline.fireChannelRead(readBuf.get(i));
    * pipeline.fireChannelReadComplete();
    * pipeline.fireExceptionCaught(exception);

---
