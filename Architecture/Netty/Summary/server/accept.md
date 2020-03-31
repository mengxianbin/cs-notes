* super(null, channel, SelectionKey.OP_ACCEPT); - AbstractNioMessageChannel / NioServerSocketChannel
* super(parent, ch, SelectionKey.OP_READ); - AbstractNioByteChannel / NioSocketChannel

---

```java
    protected AbstractNioChannel(Channel parent, SelectableChannel ch, int readInterestOp) {
        super(parent);
        this.ch = ch;
        this.readInterestOp = readInterestOp;
```
