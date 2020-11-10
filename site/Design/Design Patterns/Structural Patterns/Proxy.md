[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Design](https://mengxianbin.github.io/cs-notes/site/Design) /
[Design Patterns](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns) /
[Structural Patterns](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Structural%20Patterns) /
[Proxy](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Structural%20Patterns/Proxy)

```plantuml
@startuml

class Client {

}

interface Subject {

}

class SubjectA {

}

class Proxy {

}

Client .> Subject
SubjectA ..up|> Subject
Proxy ..up|> Subject
Proxy -> SubjectA

@enduml
```

## Examples

* AOP

---
