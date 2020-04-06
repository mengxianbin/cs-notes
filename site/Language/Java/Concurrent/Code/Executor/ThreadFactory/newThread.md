[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Language](https://mengxianbin.github.io/cs-notes/site/Language) /
[Java](https://mengxianbin.github.io/cs-notes/site/Language/Java) /
[Concurrent](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent) /
[Code](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Code) /
[Executor](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Code/Executor) /
[ThreadFactory](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Code/Executor/ThreadFactory) /
[newThread](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Code/Executor/ThreadFactory/newThread)

```java
    /**
     * Constructs a new {@code Thread}.  Implementations may also initialize
     * priority, name, daemon status, {@code ThreadGroup}, etc.
     *
     * @param r a runnable to be executed by new thread instance
     * @return constructed thread, or {@code null} if the request to
     *         create a thread is rejected
     */
    Thread newThread(Runnable r);
```
