[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Design](https://mengxianbin.github.io/cs-notes/site/Design) /
[Design Patterns](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns) /
[Creational Patterns](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Creational%20Patterns) /
[Singleton](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Creational%20Patterns/Singleton)

# Singleton

## Solution

* Simple

```plantuml
@startuml
class Singleton {
    - instance: Singleton {static}
    - Singleton()
    + getInstance(): Singleton {static}
}
@enduml
```

* Inner Class


```plantuml
@startuml
class Singleton {
    - Singleton()
    + getInstance(): Singleton {static}
}

-class InstanceHolder {
    - instance: Singleton {static}
}

Singleton +. InstanceHolder 
@enduml
```

* Enumeration

```plantuml
@startuml
enum Singleton {
    + INSTANCE {static}
}
@enduml
```

## Issues

* Thread Safe
    * volatile
    * synchronized

* Uniqueness
    * Class Load
    * Clone
    * Reflect
    * Deserialize

---
