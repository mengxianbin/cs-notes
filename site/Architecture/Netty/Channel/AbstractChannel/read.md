[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[Channel](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Channel) /
[AbstractChannel](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Channel/AbstractChannel) /
[read](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Channel/AbstractChannel/read)

```java
    @Override
    public Channel read() {
        pipeline.read();
        return this;
    }
```
