[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Pipeline](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Pipeline) /
[DefaultChannelPipeline](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Pipeline/DefaultChannelPipeline) /
[HeadContext](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Pipeline/DefaultChannelPipeline/HeadContext) /
[channelRead](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Pipeline/DefaultChannelPipeline/HeadContext/channelRead)

```java
        @Override
        public void channelRead(ChannelHandlerContext ctx, Object msg) {
            ctx.fireChannelRead(msg);
        }
```
