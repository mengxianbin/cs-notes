[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Channel](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel) /
[AbstractChannel](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel/AbstractChannel) /
[AbstractUnsafe](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel/AbstractChannel/AbstractUnsafe) /
[flush](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel/AbstractChannel/AbstractUnsafe/flush)

```java
        @Override
        public final void flush() {
            assertEventLoop();

            ChannelOutboundBuffer outboundBuffer = this.outboundBuffer;
            if (outboundBuffer == null) {
                return;
            }

            outboundBuffer.addFlush();
            flush0();
        }
```

```java
        @SuppressWarnings("deprecation")
        protected void flush0() {
            ...
            final ChannelOutboundBuffer outboundBuffer = this.outboundBuffer;
            ...
            doWrite(outboundBuffer);
            ...
        }
```
