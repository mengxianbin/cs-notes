[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Language](https://mengxianbin.github.io/cs-notes/site/Language) /
[Java](https://mengxianbin.github.io/cs-notes/site/Language/Java) /
[Concurrent](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent) /
[Summary](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Summary) /
[Monitor](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Summary/Monitor)

# Monitor

## Background

* OS semaphore

## Model

* Hasen
* Hoare
* MESA

## Issues

* Exclusion
* Synchronization

## Solution

* Waiting Queue
    * Synchronizer
    * Condition Object

## MESA `wait`

```java
lock { // synchronized or lock.lock()
    while(!condition) {
        wait();
    }
}
```

* Why lock
    * ensure condition consistent

* Why while
    * condition may has been modified by other threads who holds the lock

## When to `notify()`

* the same condition
* the same action
* need only one thread to wake up

---
