[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[EventLoop](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/EventLoop) /
[NioEventLoopGroup](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/EventLoop/NioEventLoopGroup) /
[newChild](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/EventLoop/NioEventLoopGroup/newChild)

```java
    @Override
    protected EventLoop newChild(Executor executor, Object... args) throws Exception {
        EventLoopTaskQueueFactory queueFactory = args.length == 4 ? (EventLoopTaskQueueFactory) args[3] : null;
        return new NioEventLoop(this, executor, (SelectorProvider) args[0],
            ((SelectStrategyFactory) args[1]).newSelectStrategy(), (RejectedExecutionHandler) args[2], queueFactory);
    }
```
