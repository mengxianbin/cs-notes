
```java
    @Override
    protected void doBeginRead() throws Exception {
        // Channel.read() or ChannelHandlerContext.read() was called
        final SelectionKey selectionKey = this.selectionKey;
        ...
        final int interestOps = selectionKey.interestOps();
        if ((interestOps & readInterestOp) == 0) {
            selectionKey.interestOps(interestOps | readInterestOp);
        }
    }
```

> Usage

* doBeginRead - AbstractNioChannel
    * beginRead - AbstractChannel.AbstractUnsafe
        * register0 - AbstractChannel.AbstractUnsafe
        * read - DefaultChannelPipeline.HeadContext

---