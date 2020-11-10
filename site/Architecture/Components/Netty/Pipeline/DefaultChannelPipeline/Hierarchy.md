[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Pipeline](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Pipeline) /
[DefaultChannelPipeline](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Pipeline/DefaultChannelPipeline) /
[Hierarchy](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Pipeline/DefaultChannelPipeline/Hierarchy)

```plantuml
@startuml

HeadContext -up-+ DefaultChannelPipeline
TailContext -up-+ DefaultChannelPipeline

interface ChannelInboundHandler
interface ChannelOutboundHandler

HeadContext -up-|> AbstractChannelHandlerContext
HeadContext .up.|> ChannelInboundHandler
HeadContext .up.|> ChannelOutboundHandler

TailContext -up-|> AbstractChannelHandlerContext
TailContext .up.|> ChannelInboundHandler

DefaultChannelHandlerContext -up-|> AbstractChannelHandlerContext

interface ChannelHandlerContext
interface ChannelInboundInvoker
interface ChannelOutboundInvoker

AbstractChannelHandlerContext .up.|> ChannelHandlerContext
ChannelHandlerContext -up-|> ChannelInboundInvoker
ChannelHandlerContext -up-|> ChannelOutboundInvoker

@enduml
```
