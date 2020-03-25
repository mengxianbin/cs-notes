```puml
@startuml

class Builder {
    + buildPart1()
    + buildPart2()
    + buildPart3()
    + getResult(): Product
}

class Product {

}

Builder -> Product

@enduml
```
