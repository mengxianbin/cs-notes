[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Metrics](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Metrics) /
[Metric](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Metrics/Metric) /
[Histogram](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Metrics/Metric/Histogram)

```java
/**
 * A metric which calculates the distribution of a value.
 *
 * @see <a href="http://www.johndcook.com/standard_deviation.html">Accurately computing running
 *      variance</a>
 */
public class Histogram implements Metric, Sampling, Counting {
    private final Reservoir reservoir;
    private final LongAdder count;
```
