* AbstractBootstrap.register
    * AbstractBootstrap.initAndRegister

---

* AbstractBootstrap.initAndRegister
    * ChannelFactory.newChannel
    * AbstractBootstrap.init(channel)
    * EventLoopGroup.register(channel)

---

* bind `AbstractBootstrap`
    * doBind `AbstractBootstrap`
        * initAndRegister `AbstractBootstrap`
        * doBind0 `AbstractBootstrap`
            * Channel.bind `AbstractChannel`
                * bind `DefaultChannelPipeline`
                    * DefaultChannelPipeline.tail.bind `AbstractChannelHandlerContext`
                        * findContextOutbound `AbstractChannelHandlerContext`
                        * invokeBind `AbstractChannelHandlerContext`
                            * bind `ChannelOutboundHandler` (DefaultChannelPipeline.HeadContext)
                                * bind `Channel.Unsafe` (pipeline.channel().unsafe)

---
