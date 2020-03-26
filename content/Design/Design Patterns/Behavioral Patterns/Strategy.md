[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/content) /
[Design](https://mengxianbin.github.io/cs-notes/content/Design) /
[Design Patterns](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns) /
[Behavioral Patterns](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns/Behavioral%20Patterns) /
[Strategy](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns/Behavioral%20Patterns/Strategy)

```puml
@startuml

class Context {
    - strategy: Strategy
    + handle()
}

interface Strategy {
    + handle()
}

class StrategyA {

}

class StrategyB {
    
}

Context - Strategy
Strategy <|.. StrategyA
Strategy <|.. StrategyB

@enduml
```