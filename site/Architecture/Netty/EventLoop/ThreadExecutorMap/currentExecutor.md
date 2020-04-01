[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[EventLoop](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/EventLoop) /
[ThreadExecutorMap](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/EventLoop/ThreadExecutorMap) /
[currentExecutor](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/EventLoop/ThreadExecutorMap/currentExecutor)

```java
    /**
     * Returns the current {@link EventExecutor} that uses the {@link Thread}, or {@code null} if none / unknown.
     */
    public static EventExecutor currentExecutor() {
        return mappings.get();
    }
```
