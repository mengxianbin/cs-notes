[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Metrics](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Metrics) /
[Basic](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Metrics/Basic) /
[Snapshot](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Metrics/Basic/Snapshot)

```java
    /**
     * Returns the value at the 75th percentile in the distribution.
     *
     * @return the value at the 75th percentile
     */
    public double get75thPercentile() {
        return getValue(0.75);
    }

    /**
     * Returns the value at the 95th percentile in the distribution.
     *
     * @return the value at the 95th percentile
     */
    public double get95thPercentile() {
        return getValue(0.95);
    }

    /**
     * Returns the value at the 98th percentile in the distribution.
     *
     * @return the value at the 98th percentile
     */
    public double get98thPercentile() {
        return getValue(0.98);
    }

    /**
     * Returns the value at the 99th percentile in the distribution.
     *
     * @return the value at the 99th percentile
     */
    public double get99thPercentile() {
        return getValue(0.99);
    }

    /**
     * Returns the value at the 99.9th percentile in the distribution.
     *
     * @return the value at the 99.9th percentile
     */
    public double get999thPercentile() {
        return getValue(0.999);
    }    
```

```java
    /**
     * Returns the highest value in the snapshot.
     *
     * @return the highest value
     */
    public abstract long getMax();

    /**
     * Returns the arithmetic mean of the values in the snapshot.
     *
     * @return the arithmetic mean
     */
    public abstract double getMean();

    /**
     * Returns the lowest value in the snapshot.
     *
     * @return the lowest value
     */
    public abstract long getMin();

    /**
     * Returns the standard deviation of the values in the snapshot.
     *
     * @return the standard value
     */
    public abstract double getStdDev();

    /**
     * Returns the median value in the distribution.
     *
     * @return the median value
     */
    public double getMedian() {
        return getValue(0.5);
    }
```

```java
    /**
     * Returns the value at the given quantile.
     *
     * @param quantile    a given quantile, in {@code [0..1]}
     * @return the value in the distribution at {@code quantile}
     */
    public abstract double getValue(double quantile);
```

```java
    /**
     * Returns the entire set of values in the snapshot.
     *
     * @return the entire set of values
     */
    public abstract long[] getValues();
```
