* processSeletedKeys - NioEventLoop
    * processSeletedKey - NioEventLoop
        * read - Channel.Unsafe
            * fireChannelRead - Pipeline
                * invokeChannelRead - HeadContext
                    * channelRead - ChannelHandler / HeadContext
                        * fireChannelRead - ChannelHandlerContext

---

* channelRead - ChannelHandler / TailContext
    * onUnhandledInboundMessage

---
