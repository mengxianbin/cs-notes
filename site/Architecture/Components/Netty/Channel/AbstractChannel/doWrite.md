[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Channel](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel) /
[AbstractChannel](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel/AbstractChannel) /
[doWrite](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel/AbstractChannel/doWrite)

```java
    /**
     * Flush the content of the given buffer to the remote peer.
     */
    protected abstract void doWrite(ChannelOutboundBuffer in) throws Exception;
```
