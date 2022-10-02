[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Design](https://mengxianbin.github.io/cs-notes/site/Design) /
[Design Patterns](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns) /
[Behavioral Patterns](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Behavioral%20Patterns) /
[Observer](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Behavioral%20Patterns/Observer)

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
