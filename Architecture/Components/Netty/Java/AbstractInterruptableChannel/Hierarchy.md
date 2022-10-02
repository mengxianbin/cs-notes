```plantuml
@startuml

SelectableChannel -up-|> AbstractInterruptibleChannel
AbstractSelectableChannel -up-|> SelectableChannel

SocketChannel -up-|> AbstractSelectableChannel
ServerSocketChannel -up-|> AbstractSelectableChannel

@enduml
```
