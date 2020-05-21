```java
public class Histogram implements Metric, Sampling, Counting {
    private final Reservoir reservoir;
    private final LongAdder count;
```
