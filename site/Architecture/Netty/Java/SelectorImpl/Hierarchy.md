[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[Java](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Java) /
[SelectorImpl](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Java/SelectorImpl) /
[Hierarchy](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Java/SelectorImpl/Hierarchy)

```puml
@startuml

class PollSelectorImpl <<Common>> {

}

class EPollSelectorImpl <<Linux>> {

}

class KQueueSelectorImpl <<MacOS>> {

}

AbstractSelector -up-|> Selector
SelectorImpl -up-|> AbstractSelector
PollSelectorImpl -up-|> SelectorImpl
EPollSelectorImpl -up-|> SelectorImpl
KQueueSelectorImpl -up-|> SelectorImpl

@enduml
```