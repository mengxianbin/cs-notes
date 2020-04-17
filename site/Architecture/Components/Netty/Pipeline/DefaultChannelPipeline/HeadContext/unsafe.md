[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Pipeline](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Pipeline) /
[DefaultChannelPipeline](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Pipeline/DefaultChannelPipeline) /
[HeadContext](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Pipeline/DefaultChannelPipeline/HeadContext) /
[unsafe](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Pipeline/DefaultChannelPipeline/HeadContext/unsafe)

```java
    final class HeadContext extends AbstractChannelHandlerContext
            implements ChannelOutboundHandler, ChannelInboundHandler {

        private final Unsafe unsafe;

        HeadContext(DefaultChannelPipeline pipeline) {
            super(pipeline, null, HEAD_NAME, HeadContext.class);
            unsafe = pipeline.channel().unsafe();
            setAddComplete();
        }

        @Override
        public ChannelHandler handler() {
            return this;
        }
        ...
    }
```
