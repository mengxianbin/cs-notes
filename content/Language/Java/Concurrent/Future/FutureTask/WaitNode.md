[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/content) /
[Language](https://mengxianbin.github.io/cs-notes/content/Language) /
[Java](https://mengxianbin.github.io/cs-notes/content/Language/Java) /
[Concurrent](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent) /
[Future](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent/Future) /
[FutureTask](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent/Future/FutureTask) /
[WaitNode](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent/Future/FutureTask/WaitNode)

```java
    /**
     * Simple linked list nodes to record waiting threads in a Treiber
     * stack.  See other classes such as Phaser and SynchronousQueue
     * for more detailed explanation.
     */
    static final class WaitNode {
        volatile Thread thread;
        volatile WaitNode next;
        WaitNode() { thread = Thread.currentThread(); }
    }
```