[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Bootstrap](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Bootstrap) /
[ServerBootstrap](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Bootstrap/ServerBootstrap) /
[usage](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Bootstrap/ServerBootstrap/usage)

* group `NioEventLoopGroup`
* channel `NioServerSocketChannel`
* option
    * ChannelOption.SO_BACKLOG, 128
* childOption
    * ChannelOption.SO_KEEPALIVE, true
* handler
* childHandler `ChannelInitializer`
    * handler `ChannelPipeline.addLast`
        * LengthFieldBasedFrameDecoder
        * LengthFieldPrepender
        * IdleStateHandler
        * MessageProcessor
* bind

---
