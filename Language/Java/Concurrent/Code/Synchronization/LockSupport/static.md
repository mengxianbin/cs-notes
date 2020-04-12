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
