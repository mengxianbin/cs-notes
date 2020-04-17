[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Bootstrap](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Bootstrap) /
[ServerBootstrap](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Bootstrap/ServerBootstrap) /
[register](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Bootstrap/ServerBootstrap/register)

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
