* doBind - AbstractBootstrap
    * initAndRegister - AbstractBootstrap
        * newChannel - ChannelFactory
        * init(Channel) - AbstractBootstrap
        * register(Channel channel) - EventLoopGroup / MultithreadEventLoopGroup
            * EventExecutor next() - EventLoopGroup / MultithreadEventExecutorGroup
                * EventExecutor next() - EventExecutorChooserFactory.EventExecutorChooser
            * register(Channel channel) - EventLoopGroup / SingleThreadEventLoop
                * void register(EventLoop eventLoop, ChannelPromise promise); - Channel.Unsafe

---
