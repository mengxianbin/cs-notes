[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Summary](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Summary) /
[server](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Summary/server) /
[accept](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Summary/server/accept)

* super(null, channel, SelectionKey.OP_ACCEPT); - AbstractNioMessageChannel / NioServerSocketChannel
* super(parent, ch, SelectionKey.OP_READ); - AbstractNioByteChannel / NioSocketChannel

---

```java
    protected AbstractNioChannel(Channel parent, SelectableChannel ch, int readInterestOp) {
        super(parent);
        this.ch = ch;
        this.readInterestOp = readInterestOp;
```
