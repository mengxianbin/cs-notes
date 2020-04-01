[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[Bootstrap](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Bootstrap) /
[Bootstrap](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Bootstrap/Bootstrap) /
[init](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Bootstrap/Bootstrap/init)

```java
    @Override
    @SuppressWarnings("unchecked")
    void init(Channel channel) {
        ChannelPipeline p = channel.pipeline();
        p.addLast(config.handler());

        setChannelOptions(channel, options0().entrySet().toArray(EMPTY_OPTION_ARRAY), logger);
        setAttributes(channel, attrs0().entrySet().toArray(EMPTY_ATTRIBUTE_ARRAY));
    }
```
