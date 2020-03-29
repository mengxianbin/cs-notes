[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/content) /
[Architecture](https://mengxianbin.github.io/cs-notes/content/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/content/Architecture/Netty) /
[Reactor](https://mengxianbin.github.io/cs-notes/content/Architecture/Netty/Reactor) /
[Netty](https://mengxianbin.github.io/cs-notes/content/Architecture/Netty/Reactor/Netty) /
[Implementation](https://mengxianbin.github.io/cs-notes/content/Architecture/Netty/Reactor/Netty/Implementation)

```puml
@startuml

'left to right direction

storage Client {
    (Client) as Client1
    (Client) as Client2
    (Client) as Client3

    Client1 -[hidden]- Client2
    Client2 -[hidden]- Client3
}

storage BossGroup <<Acceptor Pool>> {
    (NioEventLoopGroup) as NioEventLoopGroupBoss
}

storage WorkerGroup <<IO Thread Pool>> {
    (NioEventLoopGroup) as NioEventLoopGroupWorker
}

storage EventLoop as EventLoop1 {
    queue TaskQueue as TaskQueue1
}

storage EventLoop as EventLoop2 {
    queue TaskQueue as TaskQueue2
}

Client1 -> NioEventLoopGroupBoss
Client2 -> NioEventLoopGroupBoss
Client3 -[norank]-> NioEventLoopGroupBoss
NioEventLoopGroupBoss -> NioEventLoopGroupWorker: Dispatcher
NioEventLoopGroupWorker --> TaskQueue1
NioEventLoopGroupWorker --> TaskQueue2

component Channel as Channel1
component Channel as Channel2

TaskQueue1 --> Channel1
TaskQueue1 --> Channel2

storage Pipeline as Pipeline1 {
    [Channel Handler] as ChannelHandler1
    [Channel Handler] as ChannelHandler2
    [Channel Handler] as ChannelHandler3

    ChannelHandler1 --> ChannelHandler2
    ChannelHandler2 --> ChannelHandler3
}

storage Pipeline as Pipeline2 {

}

Channel1 --> ChannelHandler1
Channel2 --> Pipeline2

ChannelHandler1 -left- [ChannelHandlerContext]

@enduml
```
