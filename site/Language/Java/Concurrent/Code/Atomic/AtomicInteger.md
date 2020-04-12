[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Language](https://mengxianbin.github.io/cs-notes/site/Language) /
[Java](https://mengxianbin.github.io/cs-notes/site/Language/Java) /
[Concurrent](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent) /
[Code](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Code) /
[Atomic](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Code/Atomic) /
[AtomicInteger](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Code/Atomic/AtomicInteger)

```java
public class AtomicInteger extends Number implements java.io.Serializable {
    private static final long serialVersionUID = 6214790243416807050L;

    /*
     * This class intended to be implemented using VarHandles, but there
     * are unresolved cyclic startup dependencies.
     */
    private static final jdk.internal.misc.Unsafe U = jdk.internal.misc.Unsafe.getUnsafe();
    private static final long VALUE = U.objectFieldOffset(AtomicInteger.class, "value");

    private volatile int value;

    /**
     * Creates a new AtomicInteger with the given initial value.
     *
     * @param initialValue the initial value
     */
    public AtomicInteger(int initialValue) {
        value = initialValue;
    }

    /**
     * Creates a new AtomicInteger with initial value {@code 0}.
     */
    public AtomicInteger() {
    }

    /**
     * Returns the current value,
     * with memory effects as specified by {@link VarHandle#getVolatile}.
     *
     * @return the current value
     */
    public final int get() {
        return value;
    }

    ...
}
```
