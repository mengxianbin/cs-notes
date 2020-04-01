[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Design](https://mengxianbin.github.io/cs-notes/site/Design) /
[Design Patterns](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns) /
[Creational Patterns](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Creational%20Patterns) /
[Simple Factory](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Creational%20Patterns/Simple%20Factory)

```puml
@startuml

class Client {

}

class SimpleFactory {
+ createProduct(): Product
}

interface Product {

}

class ProductA {

}

class ProductB {
    
}

Client ..> SimpleFactory
Client ..> Product
SimpleFactory .> Product
SimpleFactory ..> ProductA
SimpleFactory ..> ProductB
Product <|.. ProductA
Product <|.. ProductB

@enduml
```
