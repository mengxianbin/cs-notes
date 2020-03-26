[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/content) /
[Design](https://mengxianbin.github.io/cs-notes/content/Design) /
[Design Patterns](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns) /
[Structural Patterns](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns/Structural%20Patterns) /
[Others](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns/Structural%20Patterns/Others) /
[Facade](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns/Structural%20Patterns/Others/Facade)

```puml
@startuml

class Client {

}

namespace module {

    class Facade {

    }

    ~class InternalA {

    }

    ~class InternalB {

    }

    Facade ..> InternalA
    Facade ..> InternalB
}

Client .> module.Facade

@enduml
```
