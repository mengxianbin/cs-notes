[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/content) /
[Design](https://mengxianbin.github.io/cs-notes/content/Design) /
[Design Patterns](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns) /
[Structural Patterns](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns/Structural%20Patterns) /
[Bridge](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns/Structural%20Patterns/Bridge)

```puml
@startuml

interface Consumer {
    + consume()
}

interface Supplier {
    + supply()
}

class ConsumerA {

}

class SupplierA {

}

Consumer - Supplier
ConsumerA ..up|> Consumer
SupplierA ..up|> Supplier

@enduml
```
