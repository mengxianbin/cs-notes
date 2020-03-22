[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/content) /
[Language](https://mengxianbin.github.io/cs-notes/content/Language) /
[Java](https://mengxianbin.github.io/cs-notes/content/Language/Java) /
[Concurrent](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent) /
[Executor](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent/Executor) /
[ExecutorService](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent/Executor/ExecutorService)

* shutdown
* submit
    * Runnable -> Future
    * Runnable, T -> Future
    * Callable -> Future

```java
    /**
     * Returns {@code true} if this executor has been shut down.
     *
     * @return {@code true} if this executor has been shut down
     */
    boolean isShutdown();
```

```java
    /**
     * Returns {@code true} if all tasks have completed following shut down.
     * Note that {@code isTerminated} is never {@code true} unless
     * either {@code shutdown} or {@code shutdownNow} was called first.
     *
     * @return {@code true} if all tasks have completed following shut down
     */
    boolean isTerminated();
```
