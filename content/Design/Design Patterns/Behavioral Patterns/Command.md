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

interface Invoker {

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

Client ..> CommandA
Client ..> CommandB
CommandA ..right|> Command
CommandB ..up|> Command

Client .> Invoker
Invoker ..> Command

interface Executor {

}

class Executor1 {

}

class Executor2 {

}

Invoker .> Executor
Executor1 ..up|> Executor
Executor2 ..up|> Executor

Executor ..> Command

@enduml
```
