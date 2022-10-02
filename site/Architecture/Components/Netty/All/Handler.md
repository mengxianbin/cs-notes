[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[All](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/All) /
[Handler](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/All/Handler)

```plantuml
@startuml

interface ChannelHandler {
    + handlerAdded(ChannelHandlerContext)
    + handlerRemoved(ChannelHandlerContext)
}

interface ChannelInboundHandler {
    + channelRegistered(ChannelHandlerContext)
    + channelUnregistered(ChannelHandlerContext)
    ..
    + channelActive(ChannelHandlerContext)
    + channelInactive(ChannelHandlerContext)
    ..
    + channelRead(ChannelHandlerContext, Object)
    + channelReadComplete(ChannelHandlerContext)
    + channelWritabilityChanged(ChannelHandlerContext)
    ..
    + exceptionCaught(ChannelHandlerContext, Throwable)
    + userEventTriggered(ChannelHandlerContext, Object)
}

interface ChannelOutboundHandler {
    + bind()
    + connect()
    + disconnect()
    ..
    + read()
    + write()
    + flush()
    ..
    + deregister()
    + close()
}

ChannelHandlerAdapter .up.|> ChannelHandler
ChannelInboundHandler -up-|> ChannelHandler
ChannelOutboundHandler -up-|> ChannelHandler

ChannelInboundHandlerAdapter .up.|> ChannelInboundHandler
ChannelOutboundHandlerAdapter .up.|> ChannelOutboundHandler
ChannelInboundHandlerAdapter -up-|> ChannelHandlerAdapter
ChannelOutboundHandlerAdapter -up-|> ChannelHandlerAdapter

@enduml
```
