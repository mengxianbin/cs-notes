[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Java](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Java) /
[SelectorProvider](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Java/SelectorProvider) /
[Hierarchy](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Java/SelectorProvider/Hierarchy)

```puml
@startuml

SelectorProviderImpl -up-|> SelectorProvider

PollSelectorProvider -up-|> SelectorProvider
EPollSelectorProvider -up-|> SelectorProvider
KQueueSelectorProvider -up-|> SelectorProvider

@enduml
```
