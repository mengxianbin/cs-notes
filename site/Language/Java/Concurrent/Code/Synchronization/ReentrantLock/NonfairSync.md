[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Language](https://mengxianbin.github.io/cs-notes/site/Language) /
[Java](https://mengxianbin.github.io/cs-notes/site/Language/Java) /
[Concurrent](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent) /
[Code](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Code) /
[Synchronization](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Code/Synchronization) /
[ReentrantLock](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Code/Synchronization/ReentrantLock) /
[NonfairSync](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Code/Synchronization/ReentrantLock/NonfairSync)

```java
/**
    * Sync object for non-fair locks
    */
static final class NonfairSync extends Sync {
    private static final long serialVersionUID = 7316153563782823691L;

    /**
        * Performs lock.  Try immediate barge, backing up to normal
        * acquire on failure.
        */
    final void lock() {
        if (compareAndSetState(0, 1))
            setExclusiveOwnerThread(Thread.currentThread());
        else
            acquire(1);
    }

    protected final boolean tryAcquire(int acquires) {
        return nonfairTryAcquire(acquires);
    }
}
```