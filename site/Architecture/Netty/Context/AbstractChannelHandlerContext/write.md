[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[Context](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Context) /
[AbstractChannelHandlerContext](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Context/AbstractChannelHandlerContext) /
[write](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Context/AbstractChannelHandlerContext/write)

```java
    @Override
    public ChannelFuture write(Object msg) {
        return write(msg, newPromise());
    }
```

```java
    @Override
    public ChannelFuture write(final Object msg, final ChannelPromise promise) {
        write(msg, false, promise);

        return promise;
    }
```

```java
    private void write(Object msg, boolean flush, ChannelPromise promise) {
        ...
        final AbstractChannelHandlerContext next = findContextOutbound(flush ?
                (MASK_WRITE | MASK_FLUSH) : MASK_WRITE);
        ...
        if (flush) {
            next.invokeWriteAndFlush(m, promise);
        } else {
            next.invokeWrite(m, promise);
        }
        ...                
    }
```
