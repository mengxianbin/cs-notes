[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/content) /
[Design](https://mengxianbin.github.io/cs-notes/content/Design) /
[Design Patterns](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns) /
[Structural Patterns](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns/Structural%20Patterns) /
[Bridge](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns/Structural%20Patterns/Bridge)

```puml
@startuml

class Abstract {
    - implement: Implement
    + work()
}

interface Implement {
    + handle()
}

class AbstractA {

}

class ImplementA {

}

Abstract -> Implement
AbstractA --up|> Abstract
ImplementA ..up|> Implement

@enduml
```
