[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Handler](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Handler) /
[ChannelHandler](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Handler/ChannelHandler)

```java
public interface ChannelHandler { ... }
```

```java
void handlerAdded(ChannelHandlerContext ctx) throws Exception;
```

```java
void handlerRemoved(ChannelHandlerContext ctx) throws Exception;
```

```java
    /**
     * Gets called if a {@link Throwable} was thrown.
     *
     * @deprecated if you want to handle this event you should implement {@link ChannelInboundHandler} and
     * implement the method there.
     */
    @Deprecated
    void exceptionCaught(ChannelHandlerContext ctx, Throwable cause) throws Exception;
```

---
