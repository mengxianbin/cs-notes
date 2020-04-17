[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Channel](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel) /
[Channel](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel/Channel)

```puml
@startuml

interface Channel {
    + isRegistered(): boolean
    + isOpen(): boolean
    + isActive(): boolean
    + isWritable(): boolean
    --
    + alloc(): ByteBufAllocator
    + read(): Channel
    + flush(): Channel
    --
    + id(): ChannelId
    + parent(): Channel
    + config(): ChannelConfig
    + metadata(): ChannelMetadata
    + localAddress(): SocketAddress
    + remoteAddress(): SocketAddress
    + unsafe(): Unsafe
    + eventLoop(): EventLoop
    + pipeline(): ChannelPipeline
}

interface Unsafe

Unsafe -right-+ Channel

@enduml
```
