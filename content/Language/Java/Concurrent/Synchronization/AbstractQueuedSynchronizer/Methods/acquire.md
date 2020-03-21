[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/content) /
[Language](https://mengxianbin.github.io/cs-notes/content/Language) /
[Java](https://mengxianbin.github.io/cs-notes/content/Language/Java) /
[Concurrent](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent) /
[Synchronization](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent/Synchronization) /
[AbstractQueuedSynchronizer](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent/Synchronization/AbstractQueuedSynchronizer) /
[Methods](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent/Synchronization/AbstractQueuedSynchronizer/Methods) /
[acquire](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent/Synchronization/AbstractQueuedSynchronizer/Methods/acquire)

```java
    public final void acquire(int arg) {
        if (!tryAcquire(arg) &&
            acquireQueued(addWaiter(Node.EXCLUSIVE), arg))
            selfInterrupt();
    }
```

```java
    public final void acquireInterruptibly(int arg)
            throws InterruptedException {
        if (Thread.interrupted())
            throw new InterruptedException();
        if (!tryAcquire(arg))
            doAcquireInterruptibly(arg);
    }
```

```java
    public final void acquireShared(int arg) {
        if (tryAcquireShared(arg) < 0)
            doAcquireShared(arg);
    }
```

```java
    public final void acquireSharedInterruptibly(int arg)
            throws InterruptedException {
        if (Thread.interrupted())
            throw new InterruptedException();
        if (tryAcquireShared(arg) < 0)
            doAcquireSharedInterruptibly(arg);
    }
```