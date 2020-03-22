[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/content) /
[Language](https://mengxianbin.github.io/cs-notes/content/Language) /
[Java](https://mengxianbin.github.io/cs-notes/content/Language/Java) /
[Concurrent](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent) /
[Code](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent/Code) /
[Synchronization](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent/Code/Synchronization) /
[AbstractQueuedSynchronizer](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent/Code/Synchronization/AbstractQueuedSynchronizer) /
[Methods](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent/Code/Synchronization/AbstractQueuedSynchronizer/Methods) /
[get](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent/Code/Synchronization/AbstractQueuedSynchronizer/Methods/get)


```java
    public final Collection<Thread> getWaitingThreads(ConditionObject condition) {
        if (!owns(condition))
            throw new IllegalArgumentException("Not owner");
        return condition.getWaitingThreads();
    }
```
