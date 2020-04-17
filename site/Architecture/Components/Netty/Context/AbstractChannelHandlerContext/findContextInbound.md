[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Context](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Context) /
[AbstractChannelHandlerContext](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Context/AbstractChannelHandlerContext) /
[findContextInbound](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Context/AbstractChannelHandlerContext/findContextInbound)

```java
    private AbstractChannelHandlerContext findContextInbound(int mask) {
        AbstractChannelHandlerContext ctx = this;
        do {
            ctx = ctx.next;
        } while ((ctx.executionMask & mask) == 0);
        return ctx;
    }
```
