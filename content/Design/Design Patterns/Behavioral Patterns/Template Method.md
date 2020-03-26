[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/content) /
[Design](https://mengxianbin.github.io/cs-notes/content/Design) /
[Design Patterns](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns) /
[Behavioral Patterns](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns/Behavioral%20Patterns) /
[Template Method](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns/Behavioral%20Patterns/Template%20Method)

```puml
@startuml

class Template {
    + handle()
    # handle1()
    # handle2()
    # handle3()
}

class TemplateA {
    # handle1()
    # handle3()
}

class TemplateB {
    # handle2()
    # handle3()
}

class TemplateC {
    # handle1()
    # handle2()
    # handle3()
}

Template <|-- TemplateA
Template <|-- TemplateB
Template <|-- TemplateC

@enduml
```
