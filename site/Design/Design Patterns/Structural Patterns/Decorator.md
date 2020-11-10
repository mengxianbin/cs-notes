[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Design](https://mengxianbin.github.io/cs-notes/site/Design) /
[Design Patterns](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns) /
[Structural Patterns](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Structural%20Patterns) /
[Decorator](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Structural%20Patterns/Decorator)

```plantuml
@startuml

interface Component {
    + operate()
}

class ComponentA {
    + operate()
}

interface Decorator {
    # getDecorated(): Component
    + operate()
}

ComponentA ..up|> Component
Decorator --up|> Component
Decorator *-- Component

class Decorator1 {

}

class Decorator2 {
    
}

Decorator1 ..up|> Decorator
Decorator2 ..up|> Decorator

@enduml
```
