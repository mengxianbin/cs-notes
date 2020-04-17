[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Handler](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Handler) /
[timeout](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Handler/timeout) /
[IdleStateHandler](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Handler/timeout/IdleStateHandler) /
[ReaderIdleTimeoutTask](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Handler/timeout/IdleStateHandler/ReaderIdleTimeoutTask) /
[channelIdle](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Handler/timeout/IdleStateHandler/ReaderIdleTimeoutTask/channelIdle)

```java
    /**
     * Is called when an {@link IdleStateEvent} should be fired. This implementation calls
     * {@link ChannelHandlerContext#fireUserEventTriggered(Object)}.
     */
    protected void channelIdle(ChannelHandlerContext ctx, IdleStateEvent evt) throws Exception {
        ctx.fireUserEventTriggered(evt);
    }
```
    