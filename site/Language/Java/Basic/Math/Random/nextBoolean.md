[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Language](https://mengxianbin.github.io/cs-notes/site/Language) /
[Java](https://mengxianbin.github.io/cs-notes/site/Language/Java) /
[Basic](https://mengxianbin.github.io/cs-notes/site/Language/Java/Basic) /
[Math](https://mengxianbin.github.io/cs-notes/site/Language/Java/Basic/Math) /
[Random](https://mengxianbin.github.io/cs-notes/site/Language/Java/Basic/Math/Random) /
[nextBoolean](https://mengxianbin.github.io/cs-notes/site/Language/Java/Basic/Math/Random/nextBoolean)

```java
    /**
     * Returns the next pseudorandom, uniformly distributed
     * {@code boolean} value from this random number generator's
     * sequence. The general contract of {@code nextBoolean} is that one
     * {@code boolean} value is pseudorandomly generated and returned.  The
     * values {@code true} and {@code false} are produced with
     * (approximately) equal probability.
     *
     * <p>The method {@code nextBoolean} is implemented by class {@code Random}
     * as if by:
     *  <pre> {@code
     * public boolean nextBoolean() {
     *   return next(1) != 0;
     * }}</pre>
     *
     * @return the next pseudorandom, uniformly distributed
     *         {@code boolean} value from this random number generator's
     *         sequence
     * @since 1.2
     */
    public boolean nextBoolean() {
        return next(1) != 0;
    }
```
