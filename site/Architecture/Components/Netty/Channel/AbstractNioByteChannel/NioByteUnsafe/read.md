[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Channel](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel) /
[AbstractNioByteChannel](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel/AbstractNioByteChannel) /
[NioByteUnsafe](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel/AbstractNioByteChannel/NioByteUnsafe) /
[read](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel/AbstractNioByteChannel/NioByteUnsafe/read)

* read `AbstractNioByteChannel.NioByteUnsafe`
    * final ChannelPipeline pipeline = pipeline();
    * pipeline.fireChannelRead(byteBuf);
    * pipeline.fireChannelReadComplete();
    * closeOnRead(pipeline); (if close)

---
