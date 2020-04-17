[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Handler](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Handler) /
[timeout](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Handler/timeout) /
[Hierarchy](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Handler/timeout/Hierarchy)

```puml
@startuml

ChannelInboundHandlerAdapter -up-|> ChannelHandlerAdapter
ChannelDuplexHandler -up-|> ChannelInboundHandlerAdapter
IdleStateHandler -up-|> ChannelDuplexHandler
ReadTimeoutHandler -up-|> IdleStateHandler

ChannelOutboundHandlerAdapter -up-|> ChannelHandlerAdapter
WriteTimeoutHandler -up-|> ChannelOutboundHandlerAdapter

@enduml
```

---

```puml
@startuml

AbstractIdleTask .up.|> Runnable

class AllIdleTimeoutTask
class ReaderIdleTimeoutTask
class WriterIdleTimeoutTask

ReaderIdleTimeoutTask -[hidden]- AllIdleTimeoutTask
AllIdleTimeoutTask -[hidden]- WriterIdleTimeoutTask

AllIdleTimeoutTask -right-|> AbstractIdleTask
ReaderIdleTimeoutTask -[norank]-|> AbstractIdleTask
WriterIdleTimeoutTask -[norank]-|> AbstractIdleTask

IdleStateHandler +-right- AllIdleTimeoutTask
IdleStateHandler +-[norank]- ReaderIdleTimeoutTask
IdleStateHandler +-[norank]- WriterIdleTimeoutTask

@enduml
```
