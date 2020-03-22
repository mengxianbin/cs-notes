[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/content) /
[Language](https://mengxianbin.github.io/cs-notes/content/Language) /
[Java](https://mengxianbin.github.io/cs-notes/content/Language/Java) /
[Concurrent](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent) /
[Thread](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent/Thread) /
[Fields](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent/Thread/Fields) /
[Block](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent/Thread/Fields/Block)

```java
    /**
     * The argument supplied to the current call to
     * java.util.concurrent.locks.LockSupport.park.
     * Set by (private) java.util.concurrent.locks.LockSupport.setBlocker
     * Accessed using java.util.concurrent.locks.LockSupport.getBlocker
     */
    volatile Object parkBlocker;
```

```java
    /* The object in which this thread is blocked in an interruptible I/O
     * operation, if any.  The blocker's interrupt method should be invoked
     * after setting this thread's interrupt status.
     */
    private volatile Interruptible blocker;
```

```java
    private final Object blockerLock = new Object();
```
