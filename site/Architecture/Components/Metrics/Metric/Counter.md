[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Metrics](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Metrics) /
[Metric](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Metrics/Metric) /
[Counter](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Metrics/Metric/Counter)

```java
/**
 * An incrementing and decrementing counter metric.
 */
public class Counter implements Metric, Counting {
    private final LongAdder count;

```

* Counter()
* dec()
* dec(long)
* getCount()
* inc()
* inc(long)

---
