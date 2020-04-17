[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Channel](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel) /
[AbstractChannel](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel/AbstractChannel) /
[AbstractUnsafe](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel/AbstractChannel/AbstractUnsafe) /
[write](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel/AbstractChannel/AbstractUnsafe/write)

```java
        @Override
        public final void write(Object msg, ChannelPromise promise) {
            ...
            ChannelOutboundBuffer outboundBuffer = this.outboundBuffer;
            ...
            outboundBuffer.addMessage(msg, size, promise);
        }
```
