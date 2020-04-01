[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[Pipeline](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Pipeline) /
[Hierarchy](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Pipeline/Hierarchy)

```puml
@startuml

interface ChannelPipeline
interface ChannelInboundInvoker
interface ChannelOutboundInvoker
interface Iterable

ChannelPipeline .up.|> ChannelOutboundInvoker
ChannelPipeline .up.|> Iterable
ChannelPipeline .up.|> ChannelInboundInvoker

DefaultChannelPipeline -up-|> ChannelPipeline

@enduml
```
