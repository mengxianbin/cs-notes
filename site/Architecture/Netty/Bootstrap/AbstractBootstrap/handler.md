[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[Bootstrap](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Bootstrap) /
[AbstractBootstrap](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Bootstrap/AbstractBootstrap) /
[handler](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Bootstrap/AbstractBootstrap/handler)

```java
    /**
     * the {@link ChannelHandler} to use for serving the requests.
     */
    public B handler(ChannelHandler handler) {
        this.handler = ObjectUtil.checkNotNull(handler, "handler");
        return self();
    }
```