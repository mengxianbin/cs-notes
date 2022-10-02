```plantuml
@startuml

interface Element {

}

class RealElement {

}

class NullElement {

}

Element <|.. RealElement
Element <|.. NullElement

@enduml
```
