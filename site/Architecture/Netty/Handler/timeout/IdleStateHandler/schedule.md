[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[Handler](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Handler) /
[timeout](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Handler/timeout) /
[IdleStateHandler](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Handler/timeout/IdleStateHandler) /
[schedule](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Handler/timeout/IdleStateHandler/schedule)

```java
    /**
     * This method is visible for testing!
     */
    ScheduledFuture<?> schedule(ChannelHandlerContext ctx, Runnable task, long delay, TimeUnit unit) {
        return ctx.executor().schedule(task, delay, unit);
    }
```
