[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Language](https://mengxianbin.github.io/cs-notes/site/Language) /
[Java](https://mengxianbin.github.io/cs-notes/site/Language/Java) /
[Concurrent](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent) /
[Code](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Code) /
[Synchronization](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Code/Synchronization) /
[LockSupport](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Code/Synchronization/LockSupport) /
[park](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Code/Synchronization/LockSupport/park)

```java
    public static void park() {
        UNSAFE.park(false, 0L);
    }
```

```java
    public static void park(Object blocker) {
        Thread t = Thread.currentThread();
        setBlocker(t, blocker);
        UNSAFE.park(false, 0L);
        setBlocker(t, null);
    }
```

```java
    public static void unpark(Thread thread) {
        if (thread != null)
            UNSAFE.unpark(thread);
    }
```
