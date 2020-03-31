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
                * register - SelectableChannel AbstractSelectableChannel
                    * register - AbstractSelector / SelectorImpl
                    * addKey - AbstractSelectableChannel

---
