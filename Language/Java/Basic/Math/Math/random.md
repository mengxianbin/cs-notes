```java
    /**
     * Returns a {@code double} value with a positive sign, greater
     * than or equal to {@code 0.0} and less than {@code 1.0}.
     * Returned values are chosen pseudorandomly with (approximately)
     * uniform distribution from that range.
     *
     * <p>When this method is first called, it creates a single new
     * pseudorandom-number generator, exactly as if by the expression
     *
     * <blockquote>{@code new java.util.Random()}</blockquote>
     *
     * This new pseudorandom-number generator is used thereafter for
     * all calls to this method and is used nowhere else.
     *
     * <p>This method is properly synchronized to allow correct use by
     * more than one thread. However, if many threads need to generate
     * pseudorandom numbers at a great rate, it may reduce contention
     * for each thread to have its own pseudorandom-number generator.
     *
     * @return  a pseudorandom {@code double} greater than or equal
     * to {@code 0.0} and less than {@code 1.0}.
     * @see Random#nextDouble()
     */
    public static double random() {
        return RandomNumberGeneratorHolder.randomNumberGenerator.nextDouble();
    }
```

```java
    private static final class RandomNumberGeneratorHolder {
        static final Random randomNumberGenerator = new Random();
    }
```
