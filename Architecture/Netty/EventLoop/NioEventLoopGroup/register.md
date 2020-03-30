* MultithreadEventLoopGroup

```java
    @Override
    public ChannelFuture register(Channel channel) {
        return next().register(channel);
    }
```

```java
    @Override
    public EventLoop next() {
        return (EventLoop) super.next();
    }
```

* MultithreadEventExecutorGroup

```java
    @Override
    public EventExecutor next() {
        return chooser.next();
    }
```

* SingleThreadEventLoop

```java
    @Override
    public ChannelFuture register(Channel channel) {
        return register(new DefaultChannelPromise(channel, this));
    }

    @Override
    public ChannelFuture register(final ChannelPromise promise) {
        ObjectUtil.checkNotNull(promise, "promise");
        promise.channel().unsafe().register(this, promise);
        return promise;
    }
```
