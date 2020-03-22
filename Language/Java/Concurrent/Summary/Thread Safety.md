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
