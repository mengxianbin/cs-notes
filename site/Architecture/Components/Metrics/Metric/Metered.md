[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Metrics](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Metrics) /
[Metric](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Metrics/Metric) /
[Metered](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Metrics/Metric/Metered)

```java
/**
 * An object which maintains mean and exponentially-weighted rate.
 */
public interface Metered extends Metric, Counting {
```

* getCount
* getFifteenMinuteRate
* getFiveMinuteRate
* getMeanRate
* getOneMinuteRate

---
