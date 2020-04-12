[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Language](https://mengxianbin.github.io/cs-notes/site/Language) /
[Java](https://mengxianbin.github.io/cs-notes/site/Language/Java) /
[Concurrent](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent) /
[Code](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Code) /
[Thread](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Code/Thread) /
[blocker](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Code/Thread/blocker)

```java
    /* The object in which this thread is blocked in an interruptible I/O
     * operation, if any.  The blocker's interrupt method should be invoked
     * after setting this thread's interrupt status.
     */
    private volatile Interruptible blocker;
    private final Object blockerLock = new Object();

    /* Set the blocker field; invoked via jdk.internal.access.SharedSecrets
     * from java.nio code
     */
    static void blockedOn(Interruptible b) {
        Thread me = Thread.currentThread();
        synchronized (me.blockerLock) {
            me.blocker = b;
        }
    }
```
