[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/content) /
[Design](https://mengxianbin.github.io/cs-notes/content/Design) /
[Design Patterns](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns) /
[Behavioral Patterns](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns/Behavioral%20Patterns) /
[Others](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns/Behavioral%20Patterns/Others) /
[Mediator](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns/Behavioral%20Patterns/Others/Mediator)

```puml
@startuml

interface Mediator {
    - communicators: Collection<Communicator>
    + communicate()
}

interface Communicator {
    + communicate(mediator: Mediator)
}

class MediatorA {

}

class Communicator1 {

}

class Communicator2 {
    
}

Communicator .> Mediator
Communicator -o Mediator
Communicator <|.. Communicator1
Communicator <|.. Communicator2

Mediator <|.. MediatorA

@enduml
```
