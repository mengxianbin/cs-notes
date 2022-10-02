[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Metrics](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Metrics) /
[Basic](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Metrics/Basic) /
[EWMA](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Metrics/Basic/EWMA)

```java
/**
 * An exponentially-weighted moving average.
 *
 * @see <a href="http://www.teamquest.com/pdfs/whitepaper/ldavg1.pdf">UNIX Load Average Part 1: How
 *      It Works</a>
 * @see <a href="http://www.teamquest.com/pdfs/whitepaper/ldavg2.pdf">UNIX Load Average Part 2: Not
 *      Your Average Average</a>
 * @see <a href="http://en.wikipedia.org/wiki/Moving_average#Exponential_moving_average">EMA</a>
 */
public class EWMA {
```

```java
    /**
     * Create a new EWMA with a specific smoothing constant.
     *
     * @param alpha        the smoothing constant
     * @param interval     the expected tick interval
     * @param intervalUnit the time unit of the tick interval
     */
    public EWMA(double alpha, long interval, TimeUnit intervalUnit) {
        this.interval = intervalUnit.toNanos(interval);
        this.alpha = alpha;
    }
```

```java
    private final LongAdder uncounted = new LongAdder();

    public void update(long n) {
        uncounted.add(n);
    }
```

```java
    /**
     * Mark the passage of time and decay the current rate accordingly.
     */
    public void tick() {
        final long count = uncounted.sumThenReset();
        final double instantRate = count / interval;
        if (initialized) {
            rate += (alpha * (instantRate - rate));
        } else {
            rate = instantRate;
            initialized = true;
        }
    }
```

```java
    public static EWMA oneMinuteEWMA() {
        return new EWMA(M1_ALPHA, INTERVAL, TimeUnit.SECONDS);
    }
    public static EWMA fiveMinuteEWMA() {
        return new EWMA(M5_ALPHA, INTERVAL, TimeUnit.SECONDS);
    }
    public static EWMA fifteenMinuteEWMA() {
        return new EWMA(M15_ALPHA, INTERVAL, TimeUnit.SECONDS);
    }

    private static final int INTERVAL = 5;
    private static final double SECONDS_PER_MINUTE = 60.0;
    private static final int ONE_MINUTE = 1;
    private static final int FIVE_MINUTES = 5;
    private static final int FIFTEEN_MINUTES = 15;
    private static final double M1_ALPHA = 1 - exp(-INTERVAL / SECONDS_PER_MINUTE / ONE_MINUTE);
    private static final double M5_ALPHA = 1 - exp(-INTERVAL / SECONDS_PER_MINUTE / FIVE_MINUTES);
    private static final double M15_ALPHA = 1 - exp(-INTERVAL / SECONDS_PER_MINUTE / FIFTEEN_MINUTES);
```

```java
    public double getRate(TimeUnit rateUnit) {
        return rate * (double) rateUnit.toNanos(1);
    }
```

---
