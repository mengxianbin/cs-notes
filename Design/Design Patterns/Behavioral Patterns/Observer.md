```plantuml
@startuml

interface Observable {
    + addObserver(observer: Observer)
    + removeObserver(observer: Observer)
    + notifyObservers()
}

interface Observer {
    + update(observable: Observable)
}

Observable o-> Observer

@enduml
```
