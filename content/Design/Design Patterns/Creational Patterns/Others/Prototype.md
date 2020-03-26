[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/content) /
[Design](https://mengxianbin.github.io/cs-notes/content/Design) /
[Design Patterns](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns) /
[Creational Patterns](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns/Creational%20Patterns) /
[Others](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns/Creational%20Patterns/Others) /
[Prototype](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns/Creational%20Patterns/Others/Prototype)

```puml
@startuml

class Client {

}

interface Prototype {
    + clone(): Prototype
}

class PrototypeA {
    + clone(): Prototype
}

class PrototypeB {
    + clone(): Prototype
}

Client -> Prototype
Prototype <|.. PrototypeA
Prototype <|.. PrototypeB

@enduml
```
