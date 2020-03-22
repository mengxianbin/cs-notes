[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/content) /
[Language](https://mengxianbin.github.io/cs-notes/content/Language) /
[Java](https://mengxianbin.github.io/cs-notes/content/Language/Java) /
[Concurrent](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent) /
[Code](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent/Code) /
[Synchronization](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent/Code/Synchronization) /
[AbstractQueuedSynchronizer](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent/Code/Synchronization/AbstractQueuedSynchronizer) /
[Fields](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent/Code/Synchronization/AbstractQueuedSynchronizer/Fields)

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