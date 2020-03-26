[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/content) /
[Design](https://mengxianbin.github.io/cs-notes/content/Design) /
[Design Patterns](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns) /
[Structural Patterns](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns/Structural%20Patterns) /
[Composite](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns/Structural%20Patterns/Composite)

```puml
@startuml

interface Component {
    + operate()
}

class Leaf {
    + operate()
}

class Composite {
    + operate()
}

Leaf ..up|> Component
Composite ..up|> Component
Composite o.. Component

class Client {

}

Client .> Component

@enduml
```
