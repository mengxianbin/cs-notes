[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Design](https://mengxianbin.github.io/cs-notes/site/Design) /
[Design Patterns](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns) /
[Behavioral Patterns](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Behavioral%20Patterns) /
[State](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Behavioral%20Patterns/State)

```plantuml
@startuml

class Context {
    - state: State
    ~ setState(state: State)
    ~ getState(): State
    + handle()
}

~interface State {
    + handle(context: Context)
}

class StateA {

}

class StateB {

}

Context - State
State <|.. StateA
State <|.. StateB

@enduml
```