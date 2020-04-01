[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[Channel](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Channel) /
[AbstractNioMessageChannel](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Channel/AbstractNioMessageChannel) /
[constructor](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Channel/AbstractNioMessageChannel/constructor)

```java
    /**
     * @see AbstractNioChannel#AbstractNioChannel(Channel, SelectableChannel, int)
     */
    protected AbstractNioMessageChannel(Channel parent, SelectableChannel ch, int readInterestOp) {
        super(parent, ch, readInterestOp);
    }
```
