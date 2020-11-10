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
