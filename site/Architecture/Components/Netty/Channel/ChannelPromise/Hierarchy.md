[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Channel](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel) /
[ChannelPromise](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel/ChannelPromise) /
[Hierarchy](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel/ChannelPromise/Hierarchy)

```puml
@startuml

namespace java.util.concurrent {
    interface Future
}

package io.netty {
    package util.concurrent {
        interface Promise
        interface Future

        Promise -up-|> Future
    }

    package channel {
        interface ChannelPromise
        interface ChannelFuture

        ChannelPromise -up-|> ChannelFuture
    }

    ChannelPromise -up-|> Promise
    ChannelFuture -up-|> Future

    Future  -up-|> java.util.concurrent.Future
}

@enduml
```
