[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Metrics](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Metrics) /
[Metric](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Metrics/Metric)

* *Metric*
    * Histogram
    * Counter
    * *Metered*
        * Meter
        * Timer
    * *Gauge*
        * JmxAttributeGauge
    * *MetricSet*
        * MetricRegistry
        * JvmAttributeGaugeSet

---

```java
/**
 * A tag interface to indicate that a class is a metric.
 */
public interface Metric {

}
```
