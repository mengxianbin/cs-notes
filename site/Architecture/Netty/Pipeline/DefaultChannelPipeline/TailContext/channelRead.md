[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[Pipeline](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Pipeline) /
[DefaultChannelPipeline](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Pipeline/DefaultChannelPipeline) /
[TailContext](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Pipeline/DefaultChannelPipeline/TailContext) /
[channelRead](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Pipeline/DefaultChannelPipeline/TailContext/channelRead)

```java
        @Override
        public void channelRead(ChannelHandlerContext ctx, Object msg) {
            onUnhandledInboundMessage(ctx, msg);
        }
```
