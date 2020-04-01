[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Design](https://mengxianbin.github.io/cs-notes/site/Design) /
[Design Patterns](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns) /
[Behavioral Patterns](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Behavioral%20Patterns) /
[Chain Of Responsibility](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Behavioral%20Patterns/Chain%20Of%20Responsibility)

```puml
@startuml

class Client {

}

class Request {

}

class Handler {
    - successor: Handler    
    + handle(request: Request)
}

class HandlerA {
    + handle(request: Request)
}

class HandlerB {
    + handle(request: Request)
}

Client -> Handler
Handler - Handler
Handler .> Request
Handler <|-- HandlerA
Handler <|-- HandlerB

@enduml
```

## Examples

* java.lang.ClassLoader
* DNS server - domain resolve

---
