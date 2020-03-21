[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/content) /
[Language](https://mengxianbin.github.io/cs-notes/content/Language) /
[Java](https://mengxianbin.github.io/cs-notes/content/Language/Java) /
[Concurrent](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent) /
[Synchronization](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent/Synchronization) /
[AbstractQueuedSynchronizer](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent/Synchronization/AbstractQueuedSynchronizer) /
[Node](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent/Synchronization/AbstractQueuedSynchronizer/Node)

# Fields

```java
volatile Node prev;
volatile Node next;

/**
    * The thread that enqueued this node.  Initialized on
    * construction and nulled out after use.
    */
volatile Thread thread;
```
