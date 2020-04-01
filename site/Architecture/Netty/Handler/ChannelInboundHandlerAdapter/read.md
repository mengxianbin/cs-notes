[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[Handler](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Handler) /
[ChannelInboundHandlerAdapter](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Handler/ChannelInboundHandlerAdapter) /
[read](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Handler/ChannelInboundHandlerAdapter/read)

```java
    /**
     * Calls {@link ChannelHandlerContext#fireChannelRead(Object)} to forward
     * to the next {@link ChannelInboundHandler} in the {@link ChannelPipeline}.
     *
     * Sub-classes may override this method to change behavior.
     */
    @Skip
    @Override
    public void channelRead(ChannelHandlerContext ctx, Object msg) throws Exception {
        ctx.fireChannelRead(msg);
    }
```
