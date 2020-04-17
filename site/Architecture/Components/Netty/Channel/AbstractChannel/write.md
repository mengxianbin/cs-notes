[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Channel](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel) /
[AbstractChannel](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel/AbstractChannel) /
[write](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel/AbstractChannel/write)

```java
    @Override
    public ChannelFuture write(Object msg) {
        return pipeline.write(msg);
    }
```
