[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Channel](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel) /
[AbstractNioChannel](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel/AbstractNioChannel) /
[doBeginRead](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel/AbstractNioChannel/doBeginRead)


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