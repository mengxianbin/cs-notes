[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Design](https://mengxianbin.github.io/cs-notes/site/Design) /
[Design Patterns](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns) /
[Behavioral Patterns](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Behavioral%20Patterns) /
[Command](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Behavioral%20Patterns/Command)

```plantuml
@startuml

class Client {

}

class Invoker {

}

interface Command {
    + execute()
    + undo()
    + redo()
}

class CocreteCommand {

}

Client ..> CocreteCommand
CocreteCommand ..up|> Command

Client -> Invoker
Invoker o-> Command

CocreteCommand --> CocreteReceiver

@enduml
```

## Examples

* java.util.concurrent.Executor

---
