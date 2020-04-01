[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[Summary](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Summary) /
[bootstrap](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Summary/bootstrap) /
[register](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Summary/bootstrap/register)

* boss channel
    * bind - ServerBootstrap
        * initAndRegister - ServerBootstrap
            * register - EventLoopGroup

* worker channel
    * channelRead - ServerBootstrapAcceptor
        * Channel child = msg
        * child.pipeline().addLast(childHandler) - ChannelInitializer
        * childGroup.register(child)

---

* register - AbstractChannel.AbstractUnsafe
    * AbstractChannel.this.eventLoop = eventLoop; - AbstractChannel.AbstractUnsafe
    * register0 - AbstractChannel.AbstractUnsafe
        * doRegister - AbstractChannel / AbstractNioChannel
            * selectionKey = javaChannel().register(eventLoop().unwrappedSelector(), 0, this); - AbstractNioChannel
                * register - SelectableChannel / AbstractSelectableChannel

---

* AbstractSelectableChannel

```java
public final SelectionKey register(Selector sel, int ops, Object att)
    throws ClosedChannelException
{
    ...
    SelectionKey k = findKey(sel);
    if (k != null) {
        k.attach(att);
        k.interestOps(ops);
    } else {
        // New registration
        k = ((AbstractSelector)sel).register(this, ops, att);
        addKey(k);
    }
    ...
}
```

---
