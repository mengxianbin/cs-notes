[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[EventLoop](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/EventLoop) /
[NioEventLoopGroup](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/EventLoop/NioEventLoopGroup) /
[Constructor](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/EventLoop/NioEventLoopGroup/Constructor)

```java
public class NioEventLoopGroup extends MultithreadEventLoopGroup { ... }
```

```java
    public NioEventLoopGroup() {
        this(0);
    }
```

## MultithreadEventLoopGroup

> package io.netty.channel.nio;

```java
    /**
     * @see MultithreadEventExecutorGroup#MultithreadEventExecutorGroup(int, Executor, Object...)
     */
    protected MultithreadEventLoopGroup(int nThreads, Executor executor, Object... args) {
        super(nThreads == 0 ? DEFAULT_EVENT_LOOP_THREADS : nThreads, executor, args);
    }
```

```java
    static {
        DEFAULT_EVENT_LOOP_THREADS = Math.max(1, SystemPropertyUtil.getInt(
                "io.netty.eventLoopThreads", NettyRuntime.availableProcessors() * 2));

        if (logger.isDebugEnabled()) {
            logger.debug("-Dio.netty.eventLoopThreads: {}", DEFAULT_EVENT_LOOP_THREADS);
        }
    }
```

## NettyRuntime

> package io.netty.util;

```java
    public static int availableProcessors() {
        return holder.availableProcessors();
    }
```

```java
    private static final AvailableProcessorsHolder holder = new AvailableProcessorsHolder();
```

```java
public final class NettyRuntime {
    static class AvailableProcessorsHolder { ... }
    ...
}
```

```java
        /**
         * Get the configured number of available processors. The default is {@link Runtime#availableProcessors()}.
         * This can be overridden by setting the system property "io.netty.availableProcessors" or by invoking
         * {@link #setAvailableProcessors(int)} before any calls to this method.
         *
         * @return the configured number of available processors
         */
        @SuppressForbidden(reason = "to obtain default number of available processors")
        synchronized int availableProcessors() {
            if (this.availableProcessors == 0) {
                final int availableProcessors =
                        SystemPropertyUtil.getInt(
                                "io.netty.availableProcessors",
                                Runtime.getRuntime().availableProcessors());
                setAvailableProcessors(availableProcessors);
            }
            return this.availableProcessors;
        }
```

> @see     java.lang.Runtime#getRuntime()

---
