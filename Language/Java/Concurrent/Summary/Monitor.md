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
