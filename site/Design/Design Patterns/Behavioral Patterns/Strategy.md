[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Design](https://mengxianbin.github.io/cs-notes/site/Design) /
[Design Patterns](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns) /
[Behavioral Patterns](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Behavioral%20Patterns) /
[Strategy](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Behavioral%20Patterns/Strategy)

```plantuml
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