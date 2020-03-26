[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/content) /
[Design](https://mengxianbin.github.io/cs-notes/content/Design) /
[Design Patterns](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns) /
[Structural Patterns](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns/Structural%20Patterns) /
[Decorator](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns/Structural%20Patterns/Decorator)

```puml
@startuml

class Component {
    + operate()
}

class ComponentA {
    + operate()
}

class Decorator {
    - component: Component
    + operate()
}

ComponentA --up|> Component
Decorator --up|> Component
Decorator *-- Component

@enduml
```
