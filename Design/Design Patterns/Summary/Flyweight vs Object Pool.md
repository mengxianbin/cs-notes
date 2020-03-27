## Intent

* Flyweight
    * Use sharing to support large numbers of similar objects
* Object Pool
    * Avoid expensive acquisition and release of resources by recycling objects

## Implementation

* Flyweight
    * Intern to keep each object unique and immutable
* Object Pool
    * each object may be borrowed and returned, but not unique or immutable

---
