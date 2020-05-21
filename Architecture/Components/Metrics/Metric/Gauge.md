```java
public interface Gauge<T> extends Metric {
    /**
     * Returns the metric's current value.
     *
     * @return the metric's current value
     */
    T getValue();
}
```
