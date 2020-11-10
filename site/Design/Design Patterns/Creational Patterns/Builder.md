[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Design](https://mengxianbin.github.io/cs-notes/site/Design) /
[Design Patterns](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns) /
[Creational Patterns](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Creational%20Patterns) /
[Builder](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Creational%20Patterns/Builder)

```plantuml
@startuml

class Builder {
    + setProperty1(prop1)
    + setProperty2(prop2)
    + setProperty3(prop3)
    + build(): Product
}

class Product {

}

Builder -> Product

@enduml
```
