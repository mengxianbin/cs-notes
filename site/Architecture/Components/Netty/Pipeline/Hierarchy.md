[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Pipeline](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Pipeline) /
[Hierarchy](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Pipeline/Hierarchy)

```plantuml
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
