[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Handler](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Handler) /
[ExceptionCaught](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Handler/ExceptionCaught)

## Inbound

* ChannelInboundHandler
    * exceptionCaught

## Outbound

* ChannelFuture

```java
ChannelFuture future = channel.write(message);
future.addListener(exceptionCaughtListener);
```

* ChannelPromise

```java
public class OutboundExceptionHandler extends ChannelOutboundHandlerAdapter {
    @Override
    public void write(ctx, msg, promise) {
        promise.addListener(exceptionCaughtListener);
    }
}
```

---
