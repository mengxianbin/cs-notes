[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/content) /
[Design](https://mengxianbin.github.io/cs-notes/content/Design) /
[Design Patterns](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns) /
[Behavioral Patterns](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns/Behavioral%20Patterns) /
[Command](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns/Behavioral%20Patterns/Command)

```puml
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

class CommandA {

}

class CommandB {

}

Client .> Command
Command <|.. CommandA
Command <|.. CommandB

Client .> Invoker
Invoker o-- Command

@enduml
```
