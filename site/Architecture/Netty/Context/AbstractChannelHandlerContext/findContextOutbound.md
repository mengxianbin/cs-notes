[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[Context](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Context) /
[AbstractChannelHandlerContext](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Context/AbstractChannelHandlerContext) /
[findContextOutbound](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Context/AbstractChannelHandlerContext/findContextOutbound)

```java
    private AbstractChannelHandlerContext findContextOutbound(int mask) {
        AbstractChannelHandlerContext ctx = this;
        do {
            ctx = ctx.prev;
        } while ((ctx.executionMask & mask) == 0);
        return ctx;
    }
```
