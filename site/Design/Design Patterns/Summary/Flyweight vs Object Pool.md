[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Design](https://mengxianbin.github.io/cs-notes/site/Design) /
[Design Patterns](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns) /
[Summary](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Summary) /
[Flyweight vs Object Pool](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Summary/Flyweight%20vs%20Object%20Pool)

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
