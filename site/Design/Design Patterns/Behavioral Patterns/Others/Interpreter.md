[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Design](https://mengxianbin.github.io/cs-notes/site/Design) /
[Design Patterns](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns) /
[Behavioral Patterns](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Behavioral%20Patterns) /
[Others](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Behavioral%20Patterns/Others) /
[Interpreter](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Behavioral%20Patterns/Others/Interpreter)

```plantuml
@startuml

class Client {

}

class Context {

}

interface Expression {
    + interpret(context: Context)
}

class ExpressionA {
    + interpret(context: Context)
}

class ExpressionB {
    + interpret(context: Context)
}

Client .> Context
Client .> Expression
Expression .> Context
Expression <|.. ExpressionA
Expression <|.. ExpressionB

@enduml
```
