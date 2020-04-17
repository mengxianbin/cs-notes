[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Bootstrap](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Bootstrap) /
[ServerBootstrap](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Bootstrap/ServerBootstrap) /
[ServerBootstrap](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Bootstrap/ServerBootstrap/ServerBootstrap) /
[Fields](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Bootstrap/ServerBootstrap/ServerBootstrap/Fields)

## Fields

```java
    private final Map<ChannelOption<?>, Object> childOptions = new ConcurrentHashMap<ChannelOption<?>, Object>();
    private final Map<AttributeKey<?>, Object> childAttrs = new ConcurrentHashMap<AttributeKey<?>, Object>();
    private final ServerBootstrapConfig config = new ServerBootstrapConfig(this);
    private volatile EventLoopGroup childGroup;
    private volatile ChannelHandler childHandler;
```
