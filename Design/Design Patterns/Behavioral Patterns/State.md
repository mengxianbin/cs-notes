```puml
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