[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[EventLoop](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/EventLoop) /
[ThreadExecutorMap](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/EventLoop/ThreadExecutorMap) /
[setCurrentEventExecutor](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/EventLoop/ThreadExecutorMap/setCurrentEventExecutor)

```java
private static final FastThreadLocal<EventExecutor> mappings = new FastThreadLocal<EventExecutor>();
```

```java
    /**
     * Set the current {@link EventExecutor} that is used by the {@link Thread}.
     */
    private static void setCurrentEventExecutor(EventExecutor executor) {
        mappings.set(executor);
    }
```
