```java
    private static long initialScramble(long seed) {
        return (seed ^ multiplier) & mask;
    }
```

```java
    private static final long multiplier = 0x5DEECE66DL;
    private static final long mask = (1L << 48) - 1;
```
