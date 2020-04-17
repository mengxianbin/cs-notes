* group `NioEventLoopGroup`
* channel `NioServerSocketChannel`
* option
    * ChannelOption.SO_BACKLOG, 128
* childOption
    * ChannelOption.SO_KEEPALIVE, true
* handler
* childHandler `ChannelInitializer`
    * handler `ChannelPipeline.addLast`
        * LengthFieldBasedFrameDecoder
        * LengthFieldPrepender
        * IdleStateHandler
        * MessageProcessor
* bind

---
