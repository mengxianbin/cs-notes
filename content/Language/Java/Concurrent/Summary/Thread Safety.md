[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/content) /
[Language](https://mengxianbin.github.io/cs-notes/content/Language) /
[Java](https://mengxianbin.github.io/cs-notes/content/Language/Java) /
[Concurrent](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent) /
[Summary](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent/Summary) /
[Thread Safety](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent/Summary/Thread%20Safety)

## Immutable

* final
* String
* enum
* Number
    * Integer
    * ~~AtomicInteger~~

* Collections.unmodifiableXXX()

## Exclusive Synchronization

* synchronized
* ReentrantLock

## Unblocking Synchronization

* CAS
* AtomicInteger

### ABA

## Non-synchronization

* Stack Enclosure
    * Just use local variables and arguments
    * Independent with outer public resource
    * Reentrant code (pure code)
    * Just call reentrant code
* ThreadLocal

---
