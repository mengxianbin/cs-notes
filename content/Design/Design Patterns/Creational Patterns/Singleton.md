[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/content) /
[Design](https://mengxianbin.github.io/cs-notes/content/Design) /
[Design Patterns](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns) /
[Creational Patterns](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns/Creational%20Patterns) /
[Singleton](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns/Creational%20Patterns/Singleton)

# Singleton

## Solution

* Simple

```puml
@startuml
class Singleton {
    - instance: Singleton {static}
    - Singleton()
    + getInstance(): Singleton {static}
}
@enduml
```

* Inner Class


```puml
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

```puml
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
