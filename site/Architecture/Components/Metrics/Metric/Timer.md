[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Metrics](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Metrics) /
[Metric](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Metrics/Metric) /
[Timer](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Metrics/Metric/Timer)

```java
/**
 * A timer metric which aggregates timing durations and provides duration statistics, plus
 * throughput statistics via {@link Meter}.
 */
public class Timer implements Metered, Sampling {
    /**
     * A timing context.
     *
     * @see Timer#time()
     */
    public static class Context implements Closeable {
        private final Timer timer;
        private final Clock clock;
        private final long startTime;

        private Context(Timer timer, Clock clock) {
            this.timer = timer;
            this.clock = clock;
            this.startTime = clock.getTick();
        }

        /**
         * Stops recording the elapsed time, updates the timer and returns the elapsed time in
         * nanoseconds.
         */
        public long stop() {
            final long elapsed = clock.getTick() - startTime;
            timer.update(elapsed, TimeUnit.NANOSECONDS);
            return elapsed;
        }

        @Override
        public void close() {
            stop();
        }
    }

    private final Meter meter;
    private final Histogram histogram;
    private final Clock clock;

    private void update(long duration) {
        if (duration >= 0) {
            histogram.update(duration);
            meter.mark();
        }
    }

    public void update(long duration, TimeUnit unit) {
        update(unit.toNanos(duration));
    }

    public <T> T time(Callable<T> event) throws Exception {
        final long startTime = clock.getTick();
        try {
            return event.call();
        } finally {
            update(clock.getTick() - startTime);
        }
    }

    public Context time() {
        return new Context(this, clock);
    }
```
