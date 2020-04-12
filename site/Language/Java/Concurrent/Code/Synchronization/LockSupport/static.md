[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Language](https://mengxianbin.github.io/cs-notes/site/Language) /
[Java](https://mengxianbin.github.io/cs-notes/site/Language/Java) /
[Concurrent](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent) /
[Code](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Code) /
[Synchronization](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Code/Synchronization) /
[LockSupport](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Code/Synchronization/LockSupport) /
[static](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Code/Synchronization/LockSupport/static)

```java
private static final sun.misc.Unsafe UNSAFE;

private static final long parkBlockerOffset;

static {
    try {
        UNSAFE = sun.misc.Unsafe.getUnsafe();

        Class<?> tk = Thread.class;
        parkBlockerOffset = UNSAFE.objectFieldOffset
            (tk.getDeclaredField("parkBlocker"));

    } catch (Exception ex) { throw new Error(ex); }
}
```
