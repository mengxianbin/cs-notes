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
