[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Bootstrap](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Bootstrap) /
[ServerBootstrap](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Bootstrap/ServerBootstrap) /
[ServerBootstrapAcceptor](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Bootstrap/ServerBootstrap/ServerBootstrapAcceptor) /
[ServerBootstrapAcceptor](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Bootstrap/ServerBootstrap/ServerBootstrapAcceptor/ServerBootstrapAcceptor)

```java
    private static class ServerBootstrapAcceptor extends ChannelInboundHandlerAdapter {

        private final EventLoopGroup childGroup;
        private final ChannelHandler childHandler;
        private final Entry<ChannelOption<?>, Object>[] childOptions;
        private final Entry<AttributeKey<?>, Object>[] childAttrs;
        private final Runnable enableAutoReadTask;
```
