* write - Channel
    * write - ChannelPipeline
        * write - TailContext
            * invokeWrite - ChannelHandlerContext
                * write - ChannelHandler / HeadContext
                    * write - Channel.Unsafe
                        * addMessage - ChannelOutboundBuffer

---
