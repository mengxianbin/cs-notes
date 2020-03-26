[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/content) /
[Design](https://mengxianbin.github.io/cs-notes/content/Design) /
[Design Patterns](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns) /
[Structural Patterns](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns/Structural%20Patterns) /
[Others](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns/Structural%20Patterns/Others) /
[Flyweight](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns/Structural%20Patterns/Others/Flyweight)

```puml
@startuml

class IntrinsicState {

}

class ExtrinsicState {

}

class Flyweight {
    - intrinsicState IntrinsicState
    # extrinsicState ExtrinsicState
}

class FlyweightA {
}

class FlyweightB {
}

class ExtrinsicStateA {

}

class ExtrinsicStateB {

}

Flyweight --up> IntrinsicState
Flyweight -> ExtrinsicState

ExtrinsicStateA --up|> ExtrinsicState
ExtrinsicStateB --up|> ExtrinsicState

FlyweightA --up|> Flyweight
FlyweightB --up|> Flyweight

@enduml
```
