[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Language](https://mengxianbin.github.io/cs-notes/site/Language) /
[Markdown](https://mengxianbin.github.io/cs-notes/site/Language/Markdown) /
[plantuml](https://mengxianbin.github.io/cs-notes/site/Language/Markdown/plantuml) /
[Class Diagram](https://mengxianbin.github.io/cs-notes/site/Language/Markdown/plantuml/Class%20Diagram) /
[Namespaces](https://mengxianbin.github.io/cs-notes/site/Language/Markdown/plantuml/Class%20Diagram/Namespaces)

```puml
@startuml

class BaseClass

namespace net.dummy #DDDDDD {
    .BaseClass <|-- Person
    Meeting o-- Person

    .BaseClass <|- Meeting
}

namespace net.foo {
  net.dummy.Person  <|- Person
  .BaseClass <|-- Person

  net.dummy.Meeting o-- Person
}

BaseClass <|-- net.unused.Person

@enduml# PlantUML Editor

1. select template
2. write uml diagram

@startuml

left to right direction

actor User

User --> (1. select template)
User --> (2. write uml diagram)

@enduml
```
