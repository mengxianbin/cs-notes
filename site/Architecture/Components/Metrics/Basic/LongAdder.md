[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Metrics](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Metrics) /
[Basic](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Metrics/Basic) /
[LongAdder](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Metrics/Basic/LongAdder)

```java
class LongAdder extends Striped64 implements Serializable {
```

```java
    /**
     * Equivalent to {@code add(1)}.
     */
    public void increment() {
        add(1L);
    }

    /**
     * Equivalent to {@code add(-1)}.
     */
    public void decrement() {
        add(-1L);
    }
```

```java
    /**
     * Adds the given value.
     *
     * @param x the value to add
     */
    public void add(long x) {
        Cell[] as;
        long b, v;
        HashCode hc;
        Cell a;
        int n;
        if ((as = cells) != null || !casBase(b = base, b + x)) {
            boolean uncontended = true;
            int h = (hc = threadHashCode.get()).code;
            if (as == null || (n = as.length) < 1 ||
                    (a = as[(n - 1) & h]) == null ||
                    !(uncontended = a.cas(v = a.value, v + x)))
                retryUpdate(x, hc, uncontended);
        }
    }
```

```java
    /**
     * Returns the current sum.  The returned value is <em>NOT</em> an atomic snapshot; invocation
     * in the absence of concurrent updates returns an accurate result, but concurrent updates that
     * occur while the sum is being calculated might not be incorporated.
     *
     * @return the sum
     */
    public long sum() {
        long sum = base;
        Cell[] as = cells;
        if (as != null) {
            int n = as.length;
            for (int i = 0; i < n; ++i) {
                Cell a = as[i];
                if (a != null)
                    sum += a.value;
            }
        }
        return sum;
    }

    /**
     * Resets variables maintaining the sum to zero.  This method may be a useful alternative to
     * creating a new adder, but is only effective if there are no concurrent updates.  Because this
     * method is intrinsically racy, it should only be used when it is known that no threads are
     * concurrently updating.
     */
    public void reset() {
        internalReset(0L);
    }

    /**
     * Equivalent in effect to {@link #sum} followed by {@link #reset}. This method may apply for
     * example during quiescent points between multithreaded computations.  If there are updates
     * concurrent with this method, the returned value is <em>not</em> guaranteed to be the final
     * value occurring before the reset.
     *
     * @return the sum
     */
    public long sumThenReset() {
        long sum = base;
        Cell[] as = cells;
        base = 0L;
        if (as != null) {
            int n = as.length;
            for (int i = 0; i < n; ++i) {
                Cell a = as[i];
                if (a != null) {
                    sum += a.value;
                    a.value = 0L;
                }
            }
        }
        return sum;
    }
```
