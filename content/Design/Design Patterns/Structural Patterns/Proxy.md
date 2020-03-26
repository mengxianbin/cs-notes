[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/content) /
[Design](https://mengxianbin.github.io/cs-notes/content/Design) /
[Design Patterns](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns) /
[Structural Patterns](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns/Structural%20Patterns) /
[Proxy](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns/Structural%20Patterns/Proxy)

```puml
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
