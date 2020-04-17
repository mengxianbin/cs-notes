[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Channel](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel) /
[AbstractChannel](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel/AbstractChannel) /
[AbstractUnsafe](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel/AbstractChannel/AbstractUnsafe) /
[register](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel/AbstractChannel/AbstractUnsafe/register)


```java
    private void register0(ChannelPromise promise) {
        ...
        boolean firstRegistration = neverRegistered;
        doRegister();
        neverRegistered = false;
        registered = true;

        // Ensure we call handlerAdded(...) before we actually notify the promise. This is needed as the
        // user may already fire events through the pipeline in the ChannelFutureListener.
        pipeline.invokeHandlerAddedIfNeeded();
        ...
        pipeline.fireChannelRegistered();
        ...
        if (isActive()) {
            if (firstRegistration) {
                pipeline.fireChannelActive();
            } else if (config().isAutoRead()) {
                // This channel was registered before and autoRead() is set. This means we need to begin read
                // again so that we process inbound data.
                //
                // See https://github.com/netty/netty/issues/4805
                beginRead();
            }
        }
        ...
    }
```
