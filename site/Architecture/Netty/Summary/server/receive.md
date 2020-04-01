[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[Summary](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Summary) /
[server](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Summary/server) /
[receive](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Summary/server/receive)

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
