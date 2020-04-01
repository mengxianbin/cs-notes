[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Design](https://mengxianbin.github.io/cs-notes/site/Design) /
[Design Patterns](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns) /
[Structural Patterns](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Structural%20Patterns) /
[Adapter](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Structural%20Patterns/Adapter)

```puml
@startuml

class Client {

}

class ToolA {

}

class ToolB {

}

class Adapter {
    - toolB: ToolB
}

Client .> ToolA
Adapter --up> ToolA
Adapter *-> ToolB

@enduml
```

## Examples

* java.io
    * InputStream
    * OutputStream

---
