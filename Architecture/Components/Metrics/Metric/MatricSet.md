```java
/**
 * A set of named metrics.
 *
 * @see MetricRegistry#registerAll(MetricSet)
 */
public interface MetricSet extends Metric {
    /**
     * A map of metric names to metrics.
     *
     * @return the metrics
     */
    Map<String, Metric> getMetrics();
}

```
