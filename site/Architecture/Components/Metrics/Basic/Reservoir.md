[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Metrics](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Metrics) /
[Basic](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Metrics/Basic) /
[Reservoir](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Metrics/Basic/Reservoir)

```java
package com.codahale.metrics;

/**
 * A statistically representative reservoir of a data stream.
 */
public interface Reservoir {
    /**
     * Returns the number of values recorded.
     *
     * @return the number of values recorded
     */
    int size();

    /**
     * Adds a new recorded value to the reservoir.
     *
     * @param value a new recorded value
     */
    void update(long value);

    /**
     * Returns a snapshot of the reservoir's values.
     *
     * @return a snapshot of the reservoir's values
     */
    Snapshot getSnapshot();
}

```
