[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[EventLoop](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/EventLoop) /
[NioEventLoop](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/EventLoop/NioEventLoop) /
[select](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/EventLoop/NioEventLoop/select)

```java
    private int select(long deadlineNanos) throws IOException {
        if (deadlineNanos == NONE) {
            return selector.select();
        }
        // Timeout will only be 0 if deadline is within 5 microsecs
        long timeoutMillis = deadlineToDelayNanos(deadlineNanos + 995000L) / 1000000L;
        return timeoutMillis <= 0 ? selector.selectNow() : selector.select(timeoutMillis);
    }
```
