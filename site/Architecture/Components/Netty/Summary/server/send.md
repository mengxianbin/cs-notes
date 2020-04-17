[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Summary](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Summary) /
[server](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Summary/server) /
[send](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Summary/server/send)

* write - Channel
    * write - ChannelPipeline
        * write - TailContext
            * invokeWrite - ChannelHandlerContext
                * write - ChannelHandler / HeadContext
                    * write - Channel.Unsafe
                        * addMessage - ChannelOutboundBuffer

---
