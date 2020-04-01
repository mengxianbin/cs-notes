[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Design](https://mengxianbin.github.io/cs-notes/site/Design) /
[Design Patterns](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns) /
[Behavioral Patterns](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Behavioral%20Patterns) /
[Iterator](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Behavioral%20Patterns/Iterator)

```puml
@startuml

interface Iterable<T> {
    + interator(): Iterator<T>
}

interface Iterator<T> {
    + hasNext(): boolean
    + next(): T
}

class IterableA {

}

class IteratorA {

}

Iterable .> Iterator
Iterable <|.. IterableA
Iterator <|.. IteratorA

@enduml
```
