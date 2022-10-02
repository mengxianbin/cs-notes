```plantuml
@startuml

SelectorProviderImpl -up-|> SelectorProvider

PollSelectorProvider -up-|> SelectorProvider
EPollSelectorProvider -up-|> SelectorProvider
KQueueSelectorProvider -up-|> SelectorProvider

@enduml
```
