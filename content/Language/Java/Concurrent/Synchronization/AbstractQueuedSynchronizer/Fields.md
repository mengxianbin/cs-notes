[Home](https://mengxianbin.github.io) /
[cs-note](https://mengxianbin.github.io/cs-note/content) /
[Language](https://mengxianbin.github.io/cs-note/content/Language) /
[Java](https://mengxianbin.github.io/cs-note/content/Language/Java) /
[Concurrent](https://mengxianbin.github.io/cs-note/content/Language/Java/Concurrent) /
[Synchronization](https://mengxianbin.github.io/cs-note/content/Language/Java/Concurrent/Synchronization) /
[AbstractQueuedSynchronizer](https://mengxianbin.github.io/cs-note/content/Language/Java/Concurrent/Synchronization/AbstractQueuedSynchronizer) /
[Fields](https://mengxianbin.github.io/cs-note/content/Language/Java/Concurrent/Synchronization/AbstractQueuedSynchronizer/Fields)

```java
    /**
     * Head of the wait queue, lazily initialized.  Except for
     * initialization, it is modified only via method setHead.  Note:
     * If head exists, its waitStatus is guaranteed not to be
     * CANCELLED.
     */
    private transient volatile Node head;

    /**
     * Tail of the wait queue, lazily initialized.  Modified only via
     * method enq to add new wait node.
     */
    private transient volatile Node tail;

    /**
     * The synchronization state.
     */
    private volatile int state;
```