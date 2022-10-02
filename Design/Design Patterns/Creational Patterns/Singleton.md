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
