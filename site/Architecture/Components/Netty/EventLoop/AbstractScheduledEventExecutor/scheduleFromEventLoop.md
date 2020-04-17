[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[EventLoop](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/EventLoop) /
[AbstractScheduledEventExecutor](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/EventLoop/AbstractScheduledEventExecutor) /
[scheduleFromEventLoop](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/EventLoop/AbstractScheduledEventExecutor/scheduleFromEventLoop)

```java
    final void scheduleFromEventLoop(final ScheduledFutureTask<?> task) {
        // nextTaskId a long and so there is no chance it will overflow back to 0
        scheduledTaskQueue().add(task.setId(++nextTaskId));
    }
```
