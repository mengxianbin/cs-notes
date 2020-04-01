[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Design](https://mengxianbin.github.io/cs-notes/site/Design) /
[Design Patterns](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns) /
[Structural Patterns](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Structural%20Patterns) /
[Others](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Structural%20Patterns/Others) /
[Flyweight](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Structural%20Patterns/Others/Flyweight)

```puml
@startuml

interface Flyweight {
    + operate(extrinsicState)
}

class FlyweightA {
    - intrinsicState
    + operate(extrinsicState)
}

FlyweightA ..up|> Flyweight

class FlyweightFactory {
    + getFlyweight(key): Flyweight
}

FlyweightFactory o. Flyweight
Client ..up> FlyweightFactory
Client .> FlyweightA

@enduml
```
