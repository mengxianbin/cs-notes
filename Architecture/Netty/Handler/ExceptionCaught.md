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
