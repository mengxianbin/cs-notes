[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[Channel](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Channel) /
[DefaultChannelConfig](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Channel/DefaultChannelConfig) /
[autoRead](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Channel/DefaultChannelConfig/autoRead)

```java
    private static final AtomicIntegerFieldUpdater<DefaultChannelConfig> AUTOREAD_UPDATER =
            AtomicIntegerFieldUpdater.newUpdater(DefaultChannelConfig.class, "autoRead");
    
    @SuppressWarnings("FieldMayBeFinal")
    private volatile int autoRead = 1;
```
