[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/content) /
[Design](https://mengxianbin.github.io/cs-notes/content/Design) /
[Design Patterns](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns) /
[Creational Patterns](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns/Creational%20Patterns) /
[Factory Method](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns/Creational%20Patterns/Factory%20Method)

```puml
@startuml

interface Factory {
    + createProduct(): Product
}

interface Product {

}

class ProductA {

}

class ProductB {
    
}

ProductA ..up|> Product
ProductB ...up|> Product

class FactoryA {
    + createProduct(): Product
}

class FactoryB {
    + createProduct(): Product
}

Factory -> Product
Factory <|.. FactoryA
Factory <|... FactoryB

FactoryA .> ProductA
FactoryB .> ProductB

@enduml
```
