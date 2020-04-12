[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Language](https://mengxianbin.github.io/cs-notes/site/Language) /
[Java](https://mengxianbin.github.io/cs-notes/site/Language/Java) /
[Concurrent](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent) /
[Code](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Code) /
[Thread](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Code/Thread) /
[run](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Code/Thread/run)

```java
    /**
     * If this thread was constructed using a separate
     * {@code Runnable} run object, then that
     * {@code Runnable} object's {@code run} method is called;
     * otherwise, this method does nothing and returns.
     * <p>
     * Subclasses of {@code Thread} should override this method.
     *
     * @see     #start()
     * @see     #stop()
     * @see     #Thread(ThreadGroup, Runnable, String)
     */
    @Override
    public void run() {
        if (target != null) {
            target.run();
        }
    }
```
