[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Pipeline](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Pipeline) /
[DefaultChannelPipeline](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Pipeline/DefaultChannelPipeline) /
[HeadContext](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Pipeline/DefaultChannelPipeline/HeadContext) /
[write](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Pipeline/DefaultChannelPipeline/HeadContext/write)

```java
        @Override
        public void write(ChannelHandlerContext ctx, Object msg, ChannelPromise promise) {
            unsafe.write(msg, promise);
        }
```
