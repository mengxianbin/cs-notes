[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Context](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Context) /
[AbstractChannelHandlerContext](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Context/AbstractChannelHandlerContext) /
[ordered](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Context/AbstractChannelHandlerContext/ordered)

```java
    AbstractChannelHandlerContext(DefaultChannelPipeline pipeline, EventExecutor executor,
                                  String name, Class<? extends ChannelHandler> handlerClass) {
        this.name = ObjectUtil.checkNotNull(name, "name");
        this.pipeline = pipeline;
        this.executor = executor;
        this.executionMask = mask(handlerClass);
        // Its ordered if its driven by the EventLoop or the given Executor is an instanceof OrderedEventExecutor.
        ordered = executor == null || executor instanceof OrderedEventExecutor;
    }
```
