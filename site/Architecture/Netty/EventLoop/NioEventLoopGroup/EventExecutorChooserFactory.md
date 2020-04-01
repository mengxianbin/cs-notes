[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[EventLoop](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/EventLoop) /
[NioEventLoopGroup](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/EventLoop/NioEventLoopGroup) /
[EventExecutorChooserFactory](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/EventLoop/NioEventLoopGroup/EventExecutorChooserFactory)

```java
/**
 * Factory that creates new {@link EventExecutorChooser}s.
 */
@UnstableApi
public interface EventExecutorChooserFactory {

    /**
     * Returns a new {@link EventExecutorChooser}.
     */
    EventExecutorChooser newChooser(EventExecutor[] executors);

    /**
     * Chooses the next {@link EventExecutor} to use.
     */
    @UnstableApi
    interface EventExecutorChooser {

        /**
         * Returns the new {@link EventExecutor} to use.
         */
        EventExecutor next();
    }
}

```
